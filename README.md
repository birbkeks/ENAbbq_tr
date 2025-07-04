<div align="center">
<h1>ENA: DREAM BBQ TÜRKÇE YAMA</h1>
ENA: Dream BBQ için Türkçe yama!
</div>

## İçerikler
- [Açıklama](#Açıklama)
- [Kurulum](#Kurulum)
- [Derleme](#Derleme)
- [Ekip](#Ekip)
- [License](#License)

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

@birbkeks - Tüm yazıların Türkçe'ye çevirisi

>[!NOTE]
>
> Orijinal yapımcılar @fiverebbles ve @nucl3arsnake'dir, ben sadece projeyi Rusça'dan Türkçe'ye çevirdim.

@birbkeks - ENAbbq_tr


## S.S.S
*1. Soru:* Neden tüm elementler çevirilmedi? (Ara sahnelerdeki 3D yazılar vs.)

***Cevap:*** Şuan için çevirme planımız yok - yapımcının bir görüşü ve görsel tasarımının bir parçası olduğunu düşünüyoruz.

*2. Soru:* Bende katkıda bulunmak istiyorum! Nasıl katkıda bulunabilirim?

***Cevap:*** Çeviriyi paylaşarak, hataları bildirerek ve gelişmeleri takip ederek katkıda bulunabilirsin!

## License
This project is licensed under the open license [MIT](https://mit-license.org/). You are free to use, modify and distribute this Russifier in accordance with the terms of the license.
