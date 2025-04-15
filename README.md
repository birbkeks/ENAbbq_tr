<div align="center">
<h1>TÜRKÇELEŞTİRİLMİŞ ENA: DREAM BBQ</h1>
Türkçeleştirilmiş ENA: Dream BBQ deposu
</div>

## İçerikler
- [Açıklama](#Açıklama)
- [Kurulum](#Kurulum)
- [Derleme](#Derleme)
- [Takım](#Takım)
- [Lisans](#Lisans)
- [Kanal](#Kanal)


## Açıklama
Türkçeleştirilmiş ENA: Dream BBQ tamamen manuel olarak çevirilmiştir. Neredeyse tüm oyun içi diyaloglar, arayüz ve ana menü çevirilmiştir.

Bu Türkçeleştirmede, oyunun akışına ve diğer karakterlerin bir arada yabancı dillerde konuşmasına ters olduğundan seslendirme çevirisi yer almamaktadır. Projemizin hedefi olabildiğince diğer dillerdeki yazı ve esprileri anlamını bozmadan Türkçeye çevirmek olmuştur


Kendi özel kurulum sihirbazımız güzel tasarımı, şeffaflığı ve projemizin kolaylığı için [fyne](https://fyne.io/) üzerine dayalıdır

>[!CAUTION]
>
>Türkçeleştirme beta aşamasındadır. Eğer bir hata, veya başka bir sorun ile karşılaşırsanız [Issues](https://github.com/birbkeks/ENAbbq_tr/issues) kısmından bize bildirebilirsiniz.

İyi eğlenceler!

## Kurulum
### Otomatik (Tavsiye edilir)
1. Kurulum sihirbazını [Releases](https://github.com/birbkeks/ENAbbq_tr/releases/)'den indirin.
2. Kurulum sihirbazını çalıştırın ve Türkçe yamayı kurun.
   **Not: Linux için Steam ayarlarından oyunun "Uyumluluk" seçeneğini açıp bir Proton versiyonu seçmeniz gerekmektedir**

### Manuel
1. Bütün dosyaları "resources" den indirin (SteamIcon.png ve yarnmeta hariç).
2. **resources.assets font_res.resS font_modern.resS** ve **levels** klasöründeki herşeyi ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/``` dizinine taşıyın.
3. **JoelG.ENA4.dll** ve **YarnSpinner.Unity.dll** dosyalarını ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/Managed/``` dizinine taşıyın.
4. **catalog.json** dosyasını ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/``` dizinine taşıyın.
5. **yarndialogue_assets_all_*.bundle** dosyasını ```/ENA Dream BBQ/ENA-4-DreamBBQ_Data/StreamingAssets/aa/StandaloneWindows64/``` dizinine taşıyın.

### Hardpack
**What is a hardpack?** A hardpack is a packed build of the entire game. It is needed if the migration of changes to a new game update is delayed. *It is not recommended to install hardpacks just like that. It is best to use them only in cases where the update of the Translator is delayed*
1. Go to the [Releases](https://github.com/birbkeks/ENAbbq_tr/releases/), follow the link to the TG channel with hardpacks and download the latest archive.
2. In Steam, click on the gear icon in the ENA: Dream BBQ tab, go to "Manage" and click on "Browse Local Files"
3. Unzip the contents of the archive into the game folder.

## Derleme
Eğer kurulum sihirbazını kendiniz derlemek isterseniz bu adımları takip edin:
1. Bu depoyu klonlayın:
```bash
git clone https://github.com/birbkeks/ENAbbq_tr
```
**Linux'da:**
  1. [go](https://go.dev/) ve çapraz derleme için mingw-w64-gcc indirin
  2. fyne indirin:
  ```bash
  go get fyne.io/fyne/v2@latest
  ```
  3. ``./src/`` dizinine gidin
  4. Derleyin:
  ```bash
  make linux ; make windows
  ```
  5. Çalıştırıcı dosyaları deponun ana dizinine taşıyın.

**Windows'da**
  1. [go](https://go.dev/) indirin.
  2. fyne indirin
  ```bash
  go get fyne.io/fyne/v2@latest
  ```
  3. ``./src/windows/`` dizinine gidin
  4. Windows için derleyin:
  ```bash
  go build -ldflags="-s -w -H=windowsgui" -o Installer-Windows.exe *.go
  ```
  5. Çalıştırıcı dosyaları deponun ana dizinine taşıyın.

      Not: Çapraz derleme için [bu](https://docs.fyne.io/started/cross-compiling) rehbere göz atın.

## Ekip
**Ekibimiz:**

@fiverebbles - Russifier'ın teknik kısımları, ekip koordinasyonu.

@anpatu -  Editör, karar almada yardım

@felesneveve - Rusça Çeviri, düzenleme, Telegram kanal yönetimi.

@big_fan_of_kiwi - Rusça Çeviri.

@bagaxdd - Rusça Çeviri.

@hugginggg - Rusça Çeviri.

@PhONaRr - Rusça Çeviri.

@DetectivePrince - Rusça Çeviri.

@nucl3arsnake - Rusça Çeviri, arayüzün tüm çevirisi.

@dorime_lolk4075 - Rusça Çeviri.

@kaimanDoppel - Rusça Çeviri.

@SalvetEna - Rusça Çeviri.

@birbkeks - Türkçe Çeviri.

@freeze2222 - Font.

## F.A.Q
*1. Soru:* Neden tüm elementler çevirilmedi? (Ara sahnelerdeki 3D yazılar vs.)

***Cevap:*** Şuan için çevirme planımız yok - yapımcının bir görüşü ve görsel tasarımının bir parçası olduğunu düşünüyoruz.

*2. Soru:* Bende katkıda bulunmak istiyorum! Nasıl katkıda bulunabilirim?

***Cevap:*** Çeviriyi paylaşarak, hataları bildirerek ve gelişmeleri takip ederek katkıda bulunabilirsin!

*3. Soru:* Telegram ve GitHub dışında başka bir yer var mı?

***Answer:*** Evet, [Steam Rehberi](https://steamcommunity.com/sharedfiles/filedetails/?id=3453809143) de var!

## Lisans
This project is licensed under the open license [MIT](https://mit-license.org/). You are free to use, modify and distribute this Russifier in accordance with the terms of the license.

## Channel:
<div align="center">
<img src="https://github.com/user-attachments/assets/d5718154-17b2-49a8-98be-c71cc5d5cacd" alt="image" width="335" height="470" />
</div>
