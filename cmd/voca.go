package cmd

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"lang.pkg/ent"
	"lang.pkg/ent/book"
	"lang.pkg/lib"
	"lang.pkg/router"
)

// Voca : Voca Application
type Voca struct {
	client *ent.Client
}

// Init : Init Voca Router
func (app Voca) Init() {
	router.Add(&router.CommandStruct{
		Match:  "add",
		Help:   "!add <단어장 코드>, <단어>, <뜻>, [예문]",
		Info:   "단어장에 단어를 추가하는 명령어 입니다.",
		PreRun: lib.Passport(app.client),
		Run:    app.addVoca,
	})

	router.Add(&router.CommandStruct{
		Match:  "gets",
		Help:   "!gets <단어장 코드> [형식 list/card]",
		Info:   "모든 단어를 리스트 또는 카드로 가져오는 명령어 입니다.",
		PreRun: lib.Passport(app.client),
		Run:    app.getVocas,
	})
}

func (app *Voca) addVoca(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)
	if len(args) < 3 {
		lib.CommandError(s, m, cmd)
		return
	}

	book, err := app.client.Book.
		Query().
		Where(book.BookIDEQ(args[0])).
		WithOwner().
		First(context.Background())

	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	if book.Edges.Owner.UserID != m.Author.ID {
		s.ChannelMessageSend(m.ChannelID, "💥 자신의 단어장만 편집 가능해요. 다른 사람의 단어장을 사용하고 싶으시면 `fork` 기능을 이용해주세요")
		return
	}

	entitiy := app.client.Voca.Create().SetKey(args[1]).SetValue(args[2])
	if len(args) > 3 {
		entitiy.SetExample(args[3])
	}

	voca, err := entitiy.Save(context.Background())
	if err != nil {
		log.Errorf("Failed querying voca : %v", err)
		return
	}

	book.Update().AddVocas(voca).Save(context.Background())

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("> ✅ **%s** 를 %s에 추가하였어요", voca.Key, book.Name))
}

func (app *Voca) getVocas(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)
	if len(args) < 1 {
		lib.CommandError(s, m, cmd)
		return
	}

	entity := app.client.Book.Query().
		Where(book.BookIDEQ((args[0])))

	if len(args) == 1 || len(args) > 1 && args[1] == "list" {
		book, err := entity.First(context.Background())
		if err != nil {
			log.Errorf("Failed querying book : %v", err)
			return
		}

		vocas, err := book.QueryVocas().Order(ent.Asc("created_at")).All(context.Background())
		if err != nil {
			log.Errorf("Failed querying voca : %v", err)
			return
		}

		tmpl, err := template.
			New("vocas").
			Funcs(template.FuncMap{
				"inc": func(i int) int {
					return i + 1
				},
			}).
			Parse(fmt.Sprintf("```md\n# %s / 코드: %s / 단어 %d개 \n\n{{ range $i, $v := . }}{{ inc $i }}.[{{ $v.Key }}]({{ $v.Value }}){{ with $v.Example }}\n* 예문: {{ $v.Example }}{{ end }}\n{{ end }}```", book.Name, *book.BookID, len(vocas)))

		if err != nil {
			log.Errorf("Failed parsing template : %v", err)
			return
		}

		// TODO: response가 2000자를 넘어갈경우 잘라서 페이지네이션을 구현해야 함
		var response bytes.Buffer
		if err := tmpl.Execute(&response, vocas); err != nil {
			log.Errorf("Failed executing template : %v", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, response.String())
	} else if len(args) > 1 && args[1] == "card" {
		book, err := entity.First(context.Background())
		if err != nil {
			log.Errorf("Failed querying book : %v", err)
			return
		}

		vocas, err := book.QueryVocas().Order(ent.Asc("created_at")).All(context.Background())
		if err != nil {
			log.Errorf("Failed querying voca : %v", err)
			return
		}

		embedMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "📚 " + book.Name,
			Description: "__" + book.Description + "__",
			Color:       0x70a1ff,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "공개 여부",
					Value:  map[bool]string{true: "🇴", false: "🇽"}[book.Public],
					Inline: true,
				},
				{
					Name:   "단어 개수",
					Value:  strconv.Itoa(len(vocas)),
					Inline: true,
				},
				{
					Name:   "코드",
					Value:  "`" + *book.BookID + "`",
					Inline: true,
				},
				{
					Name:  "사용 방법",
					Value: "◀️▶️ 이모지를 이용하여 단어장을 넘겨보세요.",
				},
			},
		})

		s.MessageReactionAdd(m.ChannelID, embedMsg.ID, "◀️")
		s.MessageReactionAdd(m.ChannelID, embedMsg.ID, "▶️")
		s.MessageReactionAdd(m.ChannelID, embedMsg.ID, "❌")

		var pagination int = -1

		for {
			ctx, cancle := context.WithTimeout(context.Background(), 3*time.Minute)

			reaction := lib.WaitForReaction(ctx, s, func(r *discordgo.MessageReactionAdd) bool {
				return r.UserID == m.Author.ID && embedMsg.ID == r.MessageID && r.Emoji.Name == "◀️" || r.Emoji.Name == "▶️" || r.Emoji.Name == "❌"
			})

			cancle()

			if reaction == nil || reaction.Emoji.Name == "❌" {
				break
			}

			if reaction.Emoji.Name == "◀️" {
				if pagination <= 0 {
					s.ChannelMessageEditEmbed(m.ChannelID, embedMsg.ID, &discordgo.MessageEmbed{
						Title:       "📚 " + book.Name,
						Description: "__" + book.Description + "__",
						Color:       0x70a1ff,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:   "공개 여부",
								Value:  map[bool]string{true: "🇴", false: "🇽"}[book.Public],
								Inline: true,
							},
							{
								Name:   "단어 개수",
								Value:  strconv.Itoa(len(vocas)),
								Inline: true,
							},
							{
								Name:   "코드",
								Value:  "`" + *book.BookID + "`",
								Inline: true,
							},
							{
								Name:  "사용 방법",
								Value: "◀️▶️ 이모지를 이용하여 단어장을 넘겨보세요.",
							},
						},
					})

					continue
				}

				pagination--
			} else if reaction.Emoji.Name == "▶️" {
				if pagination >= len(vocas)-1 {
					continue
				}

				pagination++
			}

			s.ChannelMessageEditEmbed(m.ChannelID, embedMsg.ID, &discordgo.MessageEmbed{
				Title: "📚 " + book.Name,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "**단어**",
						Value:  vocas[pagination].Key,
						Inline: true,
					},
					{
						Name:   "**뜻**",
						Value:  vocas[pagination].Value,
						Inline: true,
					},
				},
			})
		}

		s.MessageReactionsRemoveAll(m.ChannelID, embedMsg.ID)
	}
}
