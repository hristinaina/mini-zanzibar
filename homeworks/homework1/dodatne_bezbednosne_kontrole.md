## Uvod
U ovom odjeljku govoriće se dodatno o prethodno objašnjenim bezbjednosnim kontrolama: __hešovanje lozinki__ i __mehanizam revizije__.
Fokus će biti na njihovoj primjeni u okviru prethodno urađenog projekata i načinu na koji su one implementirane. 
Projekat na kojem će se vršiti analiza nosi naziv _Menadžer sertifikata_, a rađen je kao predmetni projekat za kurs
Informaciona bezbjednost. Odabir ovoga projekta izvršen je iz razloga što je njegov fokus bio baš na istraživanju i
implementaciji raznih bezbjednosnih kontrola i mehanizama, te ga smatramo najviše adekvatnim za upoređivanje i provjeru 
koliko dobro je to postignuto. Serverski dio aplikacije implemenitran je u programskom jeziku Java korištenjem Spring radnog okvira.

## Hešovanje lozinki

Za hešovanje lozinki u okviru projekta korišten je interfejs __Password Encoder__ iz Spring Security radnog okvira. 
Standardna implementacija interfejsa koristi __BCryptPasswordEncoder__ i nudi dvije ključne metode: 
_encode(java.lang.CharSequence rawPassword)_ i _matches(java.lang.CharSequence rawPassword, java.lang.String encodedPassword)_.
Za enkodiranje lozinke primjenjuje se heš SHA-1 (ili veća verzija) u kombinaciji sa nasumično generisanom soli ("salt") od 8 bajta (ili većom). 

Prilikom registracije, lozinku koju korisnik unese smo heširali pozivom funkcije _encode_ i dobijeni heš smo čuvali u bazi podataka. 
Prilikom login-a ili izmjene lozinke, od korisnika je traženo da unese trenutnu lozinku. Ta lozinka koju unese se poredila 
sa onom čuvanom u bazi podataka pozivom funkcije _match_, koja interno unesenu lozinku enkoduje i poredi sa hešom čuvanim u bazi.
Što se tiče prenosa lozinke preko interneta od klijentskog do serverkog dijela aplikacije, koristili smo https i podatke
smještali isključivo u tijelo zahtjeva, da bismo osigurali da je sva komunikacija enkriptovana.

Pošto se lozinke čuvaju kao heširane u bazi podataka i s obzirom da Password Encoder ne nudi mogućnost dekodovanja lozinki odnosno
ne može se dobiti izvorna (_raw_) lozinka, smatramo da je implementacija dobro odrađena i da je ispunjen zahtjev o __povjerljivosti__ lozinki.

## Mehanizam revizije

Za logovanje događaja na nivou aplikacije koristili smo __SLF4J__ (_Simple Logging Facade for Java_) radni okvir. Ovaj radni okvir je apstrakcija koja
pruža mogućnosti za implementaciju svih stavki potrebnih za dobru realizaciju mehanizma revizije.
Međutim, mnogo od toga mi u našoj aplikaciji nismo iskoristili.

U okviru aplikacije čuvali smo dva tipa logova: logove o grešci i informativne logove. Ono što nismo uradili kao po preporuci
jeste to da nismo bilježili razliku između onih logova koji predstavljaju uspješno izvršenu operaciju i informativnih logova. 
Dodatno, nismo vodili evidenciju o logovima upozorenja.

Svaki od logova daje informacije o akciji koja se desila, vremenu kada je akcija izvršena i korisniku koji je inicirao akciju, 
ako postoji. Nigdje nismo zapisivali povjerljive ili privatne informacije (poput lozinki, kreditnih kartica i slično),
a korisnike smo identifikovali putem jedinstvenog id-ja koji im je aplikacija dodijelila. Ovime je ispunjen zahtjev o 
dobroj strukturi logova pojedinačno.

Jedan od problema je što su svi logovi čuvani u istom fajlu. Obzirom da se tokom rada aplikacije generiše veliki broj logova, 
time se gubi na njihovoj čitljivosti. Nismo implementirali nikakav mehanizam pretrage ili filtriranja logova.
Bilo bi bolje da smo logove podijelili u više fajlova po nekom kriterijumu. Taj kriterijum bi mogao da bude po tipu logova 
ili karakteristici akcije koja je izvršena. Takođe bi trebalo dodati i mogućnost pretrage i filtriranja logova koji se nalaze u više fajlova.

Preporuka koju takođe nismo ispoštovali jeste pitanje dužine koliko logove čuvati u memoriji. Nismo implementirali
nikakav mehanizam koji bi nakon određenog vremena (6 mjeseci, godinu data...) vršio brisanje logova ili njihovu kompresiju.
To plaćamo bespotrebnim zauzećem memorije.

Važne osobine logova (dostupnost, neporecivost i integritet) su samo donekle ispunjene. Nismo koristili digitalne potpise da bi podržali neporecivost
zapisa i time spriječili mogućnost manipulacije sadržajem. Da bi obezbjedili integritet potrebno je da dodamo i enkripciju zapisa,
da bi onemogućili bilo kakvu izmjenu ili neovlašćen pristup.
