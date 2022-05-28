package commands

import (
	"rotom/db"

	"github.com/bwmarrin/discordgo"
)

var YardimData = discordgo.ApplicationCommand {
	Name: "yardim",
	Description: "Bot hakkında bilgi verir",
}

func YardimHandler(s *discordgo.Session, i *discordgo.InteractionCreate, u *db.UserModel) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			Content: "__**Rotom Bot**__\n\n" + 
			"Rotom sizlere Pokémon GO ve Pokémon Unite profilinizi yönetmenizde ve arkadaş bulmanızda yardımcı olacak bir bottur\n" +
			"Sunucuda sohbete başlamadan önce profil ayarlarınızı yapmanızı şiddetle tavsiye ederiz!\n\n" +
			"__**Profil Ayarlama**__\n\n" +
			"__Eğer Pokémon GO oynamıyorsanız bu adımları atlayabilirsiniz__\n" +
			"`/isim go` komutu ile Pokémon GO oyun içi isminizi ayarlayın.\n" +
			"`/seviye go` komutu ile Pokémon GO oyun içi seviyenizi ayarlayın.\n" +
			"`/kod go` komutu ile Pokémon GO arkadaşlık kodunuzu ayarlayın.\n\n" +
			"__Eğer Pokémon Unite oynamıyorsanız bu adımları atlayabilirsiniz__\n" +
			"`/isim unite` komutu ile Pokémon Unite oyun içi isminizi ayarlayın.\n" +
			"`/seviye unite` komutu ile Pokémon Unite oyun içi seviyenizi ayarlayın.\n" +
			"`/kod unite` komutu ile Pokémon GO arkadaşlık kodunuzu ayarlayın.\n\n" +
			"Profil ayarlarınızı yaptıktan sonra kendi profilinize `/profil` komutu ile bakabilir, başkalarının profiline `/profil @kullanıcı` komutu ile bakabilirsiniz.",
		},
	})
}
