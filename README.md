# 🤖 Rotom Bot

Rotom sizlere Pokémon GO ve Pokémon Unite profilinizi yönetmenizde ve arkadaş bulmanızda yardımcı olacak bir bottur
Sunucuda sohbete başlamadan önce profil ayarlarınızı yapmanızı şiddetle tavsiye ederiz!

# 🔗 Önemli Bağlantılar
➽ Pokémon Kafé Davet Linki: https://338.rocks/pokemon

# 📜 Profil Ayarlama

Eğer Pokémon GO oynamıyorsanız bu adımları atlayabilirsiniz
- `/isim go` komutu ile Pokémon GO oyun içi isminizi ayarlayın.
- `/seviye go` komutu ile Pokémon GO oyun içi seviyenizi ayarlayın.
- `/kod go` komutu ile Pokémon GO arkadaşlık kodunuzu ayarlayın.

Eğer Pokémon Unite oynamıyorsanız bu adımları atlayabilirsiniz
- `/isim unite` komutu ile Pokémon Unite oyun içi isminizi ayarlayın.
- `/seviye unite` komutu ile Pokémon Unite oyun içi seviyenizi ayarlayın.
- `/kod unite` komutu ile Pokémon GO arkadaşlık kodunuzu ayarlayın.

Profil ayarlarınızı yaptıktan sonra kendi profilinize `/profil` komutu ile bakabilir, başkalarının profiline `/profil @kullanıcı` komutu ile bakabilirsiniz.

# 🙌 Katkıda bulunma
Bot ile ilgili aklınızdaki yenilikleri bir issue açarak belirtebilir veya kendiniz ekleyip bir pull request açabilirsiniz.
- `.env.example` dosyasını `.env` olarak yeniden adlandırın ve içerisini doldurun.
- `go mod tidy` komutu ile gerekli paketleri indirin.
- `go run . --all` komutu ile database kurulumunu yapın ve komutları Discord API'ye pushlayın.
- `go run .` komutu ile botu başlatın.
- Yeni bir komut eklediyseniz `go run . --commands` komutunu çalıştırın.
- Databasede bir değişiklik yaptıysanız `go run . --db` komutunu çalıştırın.
- Bir pull request açın.
