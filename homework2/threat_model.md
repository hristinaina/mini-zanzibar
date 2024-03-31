## ZADATAK 1
1. __Konkurencija u industirji__: Konkurentne turističke kompanije ili agencije mogu biti zainteresovane za napade kako bi oslabile reputaciju ili funkcionalnost MegaTravel platforme. Njihovi napadi mogu biti usmereni na ometanje usluga, krađu intelektualne svojine, sabotažu sistema i slično. Nivo vještine može biti visok, jer imaju stručnost i resurse za ciljanje specifičnih ranjivost ili izvršavanje složenih napada.
2. __Insajderi__ (zaposleni ili bivši zaposleni MegaTravela): ukoliko imaju nezadovoljstvo ili motivaciju da oštete kompaniju. Insajderi takođe mogu biti motivisani ličnom dobiti, kao što su krađa kreditnih kartica i drugih povjerljivih podataka. Imaju visok nivo znanja o sistemu, što ih čini posebno opasnim. Njihovi ciljevi mogu uključivati i sabotažu sistema ili štetu reputaciji kompanije. Insajderi koji djeluju samostalno  mogu da predstavljaju niži nivo opasnosti, međutim, udruženje njihovog unutrašnjeg poznavanja sistema sa iskusnim hakerima može da napravi veliku štetu po softverski sistem kompanije.
3. __Nezadvoljni korisnici__: motivacija je izražavanje nezadovoljstva i sabotiranje usluga. Nivo vještine može biti različit, od nižeg do visokog, u zavisnosti od njihovih tehničkih sposobnosti ili angažovanja drugih napadača. Mogu koristiti jednostavnije tehnike kao što su brute force napadi i manipulacija korisničkim interfejsom. Krajnji cilj može biti bilo kakva vrsta oštećenja kompanije zbog koje se pojavilo nezadovoljstvo.
4. __Kriminalne grupe__: mogu pokušati izvršiti različite vrste napada radi finansijske dobiti poput krađe finansijskih informacija (credit card information), ličnih podataka, ucjena i drugo. Mogu koristiti različite tehnike kao što su phishing, malware, DDoS napadi ili zero-day ranjivosti. Kriminalne grupe predstavljaju visok nivo opasnosti i vjerovatno najopasniju grupu napadača od gorepomenutih.
5. __Hakovanje iz zabave__: pojedinci ili grupe koje napadaju sisteme iz zabave, bez specifičnih motiva ili ciljeva. Ovi napadi se obično zovu ,,napadi iz zabave’’ ili ,,haktivizam (hactivity). Mogu predstavljati malu, ali i visoku prijetnju, u zavisnosti kakve vještine posjeduju. Motivacija takođe može biti vježba ili sticanje iskustva kako bi se kasnije mogli postići ozbiljniji ciljevi od hakovanje MegaTravel-a.
6. __Hacktivisti__: osobe ili grupe koje koriste svoje tehničke vještine u cilju podrške socijalnim, političkim ili aktivističkim ciljevima. Mogu izvršiti napade kako bi privukli pažnju javnosti na određene probleme ili kako bi protestovali protiv nečega. Obično ne predstavljaju veliku prijetnju za softverski sistem.
7. __Ekološki aktivisti__: grupacije ili pojedinci koji se bave zaštitom životne sredine. Motivacija je sabotaža MegaTravel-a zbog negativnog uticaja po životnu sredinu, kao što su avionski letovi.
Ne predstavljaju visoku opasnost po sistem. 
8. __Podzemne grupe__: kriminalne organizacije koje se bave ilegalnim aktivnostima poput krijumčarenja, pranja novca ili trgovine opojnih sredstava. Motivacija je mogućnost korišćenja MegaTravel-a radi ispunjenja ciljeva. Nivo opasnosti ove grupe napadača može da varira u zavisnosti od nivoa hakerskog znanja i motivacije.

## ZADATAK 2
### 1. Korisnički podaci: 
- __Inherentna izloženost__: Korisnički podaci imaju visoku inherentnu izloženost jer su neophodni za pružanje personalizovanih usluga putovanja. Pristup imovini imaju zaposleni koji rade na obradi rezervacija i planiranju putovanja, kao i korisnici koji pristupaju platformi za rezervaciju i upravljanje svojim putovanjima.
- __Bezbednosni ciljevi__: Poverljivost (osiguravanje privatnosti podataka klijenata), integritet (osiguravanje tačnosti podataka - podaci nisu izmenjeni ili oštećeni), dostupnost (osiguravanje pristupa ličnim informacijama poput imena, adresa, brojeva telefona, e-mail adresa).
- __Uticaj oštećenja__: Oštećenje poverljivosti može dovesti do gubitka poverenja klijenata, kršenja zakona o privatnosti podataka, gubitka poslovnog statusa. Oštećenje integriteta podataka može dovesti do pogrešnih rezervacija, gubitka podataka o putnicima, finansijskih grešaka. Oštećenje dostupnosti može dovesti do nemogućnosti pristupa podacima o putovanjima, kašnjenja u rezervacijama, gubitka poslovnih prilika.

### 2. Finansijski podaci:
- __Inherentna izloženost__: Finansijski podaci, uključujući informacije o kreditnim karticama i bankovnim računima, imaju visoku inherentnu izloženost jer su kritični za procesiranje plaćanja i fakturisanje usluga putovanja. Pristup ovim informacijama imaju finansijski timovi, računovodstvo, menadžeri prihoda.
- __Bezbednosni ciljevi__: Poverljivost (osiguravanje tajnosti finansijskih informacija - transakcija i ličnih finansijskih informacija korisnika), integritet (osiguravanje tačnosti finansijskih izveštaja), dostupnost (osiguravanje pristupa finansijskim podacima kada je to potrebno)
- __Uticaj oštećenja__: Oštećenje poverljivosti može dovesti do krađe identiteta, finansijske prevare, gubitka poverenja klijenata. Oštećenje integriteta može dovesti do finansijskih grešaka, neispravnih izveštaja, poremećaja u poslovnim transakcijama. Oštećenje dostupnosti može dovesti do kašnjenja u finansijskim transakcijama, problema sa fakturisanjem, gubitka poslovnih prilika.

### 3. Podaci o rezervacijama:
- __Inherentna izloženost__: rizik od lažnih rezervacija, manipulacija podacima rezervacija, nedostatak autentičnosti podataka
- __Bezbjednosti ciljevi__: verifikacija autentičnosti rezervacija, osiguranje intergriteta podataka i sprečavanje manipulacija podacima. To možemo postići autentifikacijom i autorizacijom, enkripcijom podataka, validacijom unosa (npr. sprečavanje XSS napada)...
- __Uticaj oštećenja__: Oštećenje ovih bezbednosnih ciljeva može dovesti do netačnih rezervacija, gubitka podataka o putnicima i financijskih gubitaka usled neslaganja u evidenciji, kao i gubitka reputacije kompanije

### 4. Interni poslovni procesi i dokumentacija:
- __Inherentna izloženost__: Menadžment, administrativni timovi, ljudski resursi.
- __Bezbednosni ciljevi__: Poverljivost (zaštita internih procedura i informacija), integritet (tačnost i autentičnost internih dokumenata), dostupnost (pristup dokumentaciji kada je to potrebno).
- __Uticaj oštećenja__: Oštećenje poverljivosti može dovesti do gubitka konkurentske prednosti, izloženosti internih procedura konkurentima, kršenja zakona o zaštiti poslovnih tajni. Oštećenje integriteta može dovesti do neispravnih odluka menadžmenta, pogrešnih internih procedura, gubitka efikasnosti. Oštećenje dostupnosti može dovesti do kašnjenja u radu, nesmetanog obavljanja poslovnih procesa, gubitka produktivnosti.

### 5. Infrastruktura i tehnološki resursi:
- __Inherentna izloženost__: Pristup imaju: IT osoblje, operativni timovi, menadžment. Tehnička infrastruktura kompanije, uključujući servere, mrežnu opremu i softverske aplikacije, ima visoku inherentnu izloženost jer su ključni za pružanje usluga putovanja i čuvanje podataka.
- __Bezbednosni ciljevi__: Integritet i dostupnost tehnoloških resursa su ključni za održavanje funkcionalnosti platforme i usluga putovanja. Poverljivost tehničkih podataka (kao što su konfiguracije servera i softverskih aplikacija) takođe je važna kako bi se sprečilo neovlašćeno pristupanje i manipulacija.
- __Uticaj oštećenja__: Oštećenje infrastrukture i tehnoloških resursa može dovesti do prekida u pružanju usluga putovanja, gubitka poslovnih prilika i ugleda, kao i potencijalnih finansijskih gubitaka usled troškova popravki i nadoknade štete.


## ZADATAK 3
### 1. Klijenti/Korisnici:
- __Interakcija__: Korisnici pristupaju MegaTravel platformi kako bi pretraživali i rezervisali smještaj, prevoz, putovanje i ostale mogućnosti i funkcionalnosti koje aplikacija nudi. Mogu da pristupe putem:
  - veb stranice,
  - mobilne aplikacije 	
- __Potencijalne ulazne tačke za napad__:
  - Korisnički interfejs: Nebezbjedno upravljanje sesijama, XSS (Cross-Site Scripting) ranjivosti, CSRF (Cross-Site Request Forgery) napadi.
  - Forme za unos podataka: SQL injection, nepotpuna validacija inputa, XML injection.
  - Autentikacija i autorizacija: Slabe lozinke, neadekvatna upotreba sesija, nedostatak dvostrukog faktora autentikacije.
### 2. Administratori sistema:
- __Interakcija__: Administratori imaju privilegovan pristup različitim dijelovima MegaTravel platforme, uključujući administratorski panel, bazu podataka, servere i druge resurse. Pristupaju administratorskom panelu kako bi upravljali sistemom, korisnicima, rezervacijama i drugim administrativnim zadacima.
- __Potencijalne ulazne tačke za napad__:
  - Administratorski panel: Slabe administratorske lozinke, nedostatak ograničenja pristupa, neadekvatno logovanje i praćenje aktivnosti.
  - Interni alati: Ranjivosti u softverskim alatima za upravljanje, neadekvatno upravljanje privilegijama, nedostatak autentifikacije.
### 3. Eksterni sistemi/partneri:
- __Interakcija__:  MegaTravel može imati integrisane sisteme sa partnerima poput: 
  - plaćanje putem PayPal-a
  - hotelski rezervacioni sistemi
  - avio-kompanije
  - agencije za iznajmljivanje automobila
  - restorani i ugostiteljski objekti
- __Potencijalne ulazne tačke za napad__:
  - API integracije: Nebezbedne API endpointe, nevalidirani ulazni podaci, nedostatak autentifikacije i autorizacije.
  - Spoljni servisi: Ranjivosti u spoljnim servisima koje MegaTravel koristi
### 4. Automatizovani sistemi/bots:
- __Interakcija__:  Automatizovani sistemi mogu se koristiti za automatizovanu pretragu, rezervaciju putovanja, praćenje cena i druge zadatke.
- __Potencijalne ulazne tačke za napad__:
  - API endpointi: Slabe kontrole pristupa, preopterećenje API-ja, zloupotreba API endpointa.
  - Automatizovani alati: Ranjivosti u skriptama i algoritmima koji se koriste za automatizaciju, nedostatak zaštite od zloupotrebe. Na primjer korisnik bi mogao kroz komunikaciju sa botom da ga navede da mu izvrši rezervaciju putovanja i da pritom izbjegne plaćanje.
 
## ZADATAK 4
![threat_model](https://github.com/hristinaina/secure-software-engineering/assets/96604086/f368550f-0f8a-4da3-851d-1172c0168250)

## ZADATAK 5


### S T R I D E (Spoofing, Tampering, Repudiation, Information disclosure, Denial of service and Elevation of privilege)

### S opasnosti:
1. lažno predstavljanje korisnika
2. mogućnost zaobilaska autentifikacije
3. man-in-the-middle slabost autentifikacije
4. brute force otkrivanje lozinki
5. skladištenje lozinki
6. nebezbjedna podrazumijevana lozinka admina
7. nedostatak vremenskog isteka sesije

### S ublažavanja opasnosti:
1. uvođenje 2FA i/ili biometrijska autentifikacija, digitalni potpisi, sigurnosni tokeni
2. OAuth ili sličan prtokol autentifikacije
3. HTTPS komunikacija, digitalni sertifikati i sigurnosni tokeni
4. postavljanje ograničenja za broj pokušaja prijavljivanja, zahtjevanje kreiranja ,,jakih’’ lozinki (kombinacija slova, brojeva i specijalnih znakova)
5. koristiti sigurne algoritme za heširanje
6. zahtjevati promjenu podrazumijevane lozinke nakon prvog prijavljivanja na sistem, postaviti snažne lozinke i zabraniti upotrebu lakih lozinki i prepoznatljivih šablona
7. postaviti odgovarajuće vremenske limite za sesije kako bi se automatski završavale nakon određenog perioda neaktivnosti

### T opasnosti:
1. malver (malicious software) manipuliše korisnikovim unosom, što može dovesti do slanja nevažećih ili zlonamernih podataka serveru.
2. malver menja poruke koje se šalju između klijenta i servera
3. malver prikazuje lažne ili zlonamerne podatke na korisničkom interfejsu
4. malver modifikuje unos sa korisnikove tastature 
### T ublažavanje opasnosti:
1. implementacija mehanizama za validaciju korisničkog unosa na serveru kako bi se sprečilo slanje nevažećih ili zlonamernih podataka.
2. korišćenje https komunikaciju, kako bi se osigurala tajnost i integritet podataka
3. redovno ažuriranje softver
4. korišćenje https protokola kako bi se sprečilo presretanje korisničkog unosa

### R opasnosti: 
1. korisnik poriče određenu transakciju
2. nedostatak mehanizma digitalnih potpisa
3. loši log zapisi
4. nedostatak log zapisa
5. neadekvatni log zapisi
### R ublažavanje opasnosti:
1. implementacija mehanizma za bilježenje i čuvanje logova transakcija, uključujući korisničke aktivnosti, vremenske odrednice i dodatne relevantne informacije
2. implementacija digitalnih potpisa za važne transakcije i korisničke akcije kako bi postigli autentičnost i integritet podataka 
3. definisati jake standarde za vođenje logova (prethodni domaći)
4. definisati jake mehanizme zaštite log podataka (prethodni domaći)
5. definisati strukturirane formate log zapisa koji sadrže relevantne informacije kao što su ko je izvršio akciju, opis akcije, vremenska odrednica, odgovor sistema  i dodatne informacije koje su neophodne (prethodni domaći)

### I opasnosti:
1. phishing korisnika
2. nezaštićen prenos podataka
3. neovlašćen pristup podacima
4. loše konfigurisane dozvole za pristup
5. nedostatak enkripcije podataka
6. slabe politike bezbjednosti
### I ublažavanje opasnosti:
1. edukacija korisnika o phishingu, verifikacija identiteta, sigurnosna obavještenja o prijetnjama ribarenja i upozoravanje na sumnjive aktivnosti ili poruke koje su primili, instalacija bezbjednog softvera koji može da detektuje phishing pokušaje
2. HTTS, SSL/TLS za enkripciju podataka tokom prenosa
3. implementacija autorizacije i autentifikacije kako bi se ograničio pristup osjetljivim informacijama samo ovlašćenim korisnicima
4. redovno ažuriranje sistema i primjena bezbjednosnih zakrpa radi otklanjanja ranjivosti koje bi mogle dovesti do neovlašćenog pristupa informacijama
5. korišćenje enkripcije podataka čak i na odmorima u sistemu (vremenski periodi kada se neka funkcija ili aktivnost ne izvršava ili je suspendovana)
6. definisanje jasnih politika i procedura za zaštitu informacijaka

### D opasnosti:
1. preopterećenje kanala, što dovodi do zastoja ili pada obrade logova
2. nepravilni parametri dovode do velikog broja korišćenja memorije servera ili CPU-a
3. nepravilni parametri dovode do zastoja ili pada servera
4. višestruke istovremene operacije dovode do nepostojanja odgovora servera
5. veliki broj mrežnih paketa blokira mrežu
6. ugrožavanje integriteta poruke


### D ublaževanje opasnosti:
1. implementiranje ograničavanje broja istovremenih zahteva, korišćenje keširanja
2. validacija unosa kako bi se sprečilo slanje nevažećih ili zlonamernih parametara. Implementacija mehanizama za ograničavanje resursa kao što su limiti memorije i CPU-a
3. korišćenje softvera koji može da se nosi sa neočekivanim ulazima i situacijama. Redovno ažuriranje softvera radi ispravljanja poznatih ranjivosti.
4. korišćenje redova čekanja ili ograničavanje broja istovremenih veza
5. korišćenje softvera koji mogu da detektuju i odbiju DDoS napade. Implementacija strategija za detekciju i filtriranje sumnjivog mrežnog saobraćaja
6. implementacija digitalnih potpisa i korišćenje sigurnih protokola komunikacije, kao što su HTTPS ili SSL/TLS.

### E opasnosti:
1. ranjivosti kernela drajvera klijentskog sistema
2. ranjivosti komponenata na klijentskoj strani
3. osetljivost na cross-domain napade
4. osetljivost na XSS napade


### E ublaževanje opasnosti:
1. redovno ažuriranje sistema i drajvera
2. redovno ažuriranje aplikacija i softverskih komponenti na klijentskoj strani radi uvođenja sigurnosnih ispravki
3. pravilno konfigurisanje CORS (Cross-Origin Resource Sharing) politika kako bi se ograničio pristup resursima sa drugih domena
4. implementacija input validacije kako bi se sprečilo ubacivanje skripti u korisnički unos. Korišćenje bezbednih API-ja i metoda za manipulaciju DOM-om kako bi se osiguralo bezbedno renderovanje korisničkog unosa.



