
# Uber (backend)

## A. Opis projekta
Uber je mobilna i web aplikacija implemenitrana po ugledu na postojeće aplikacije (Uber, Wolt, Lift, ...) koje pružaju korisnicima mogućnost da nađu prevoz i izvrše uplatu preko interneta. Ideja je da se maksimalno olakša trasnsport korisnika uz redukciju interakcije sa prevoznikom kako bi se cijeli proces ubrzao, bio konzistentniji i sigurniji. 

Postoje 4 tipa korisnika aplikacije: neregistorvani korisnik, registrovani korisnik, vozač i admin.

Aplikacija pruža neregistrovanim korisnicima mogućnost registracije, logina ili praćenja stanja zauzetosti svih vozila, što se prati preko mape. Registrovani korisnici mogu da naruče vozilo/prevoz (pri čemu navode parametre poput tipa vozila, prevoza kućnih ljubimaca, bebe u autu, itd) ili da izvrše rezervaciju istog. U realnom vremenu mogu da prate kretanje vozila i da šalju poruke svom vozaču ili grupi sa kojom se voze. Nakon završetka vožnje, mogu da ostave ocjenu i komentar vozaču ili vozilu. Pored navedenih osnovnih funkcionalnosti, korisnik takođe može da vidi istoriju svih svojih akcija na sistemu, chat-ove sa ostalim korisnicima, kao i statistiku o finansijama. 

Projekat se sastoji iz 3 komponente: 
- backend - implementiran u programskom jeziku Java, korištenjem Spring Boot radnog okvira
- frontend - implementiran u programskom jeziku TypeScript, korištenjem Angular radnog okvira
- android - implementiran u programskom jeziku Java

U ovom radu davaće se revizija koda implementiranog za __backend__ komponentu sistema.

## B. Lista članova razvojnog tima
- Anastasija Savić SV7/2020
- Katarina Vučić SV29/2020
- Hristina Adamović SV32/2020

## C. Opis pronađenih defekata
### 1. Otkrivanje informacija preko stack trace-a
> @Override
    public void commence(HttpServletRequest request, HttpServletResponse response, AuthenticationException authException)
            throws IOException {
        response.sendError(HttpServletResponse.SC_UNAUTHORIZED, **authException.getMessage()**);    

Programeri često dodaju stack trace u poruke o grešci, kao pomoć za otklanjanje grešaka. 
Nažalost, iste informacije mogu biti korisne napadaču. Niz imena klasa u praćenju steka može otkriti strukturu aplikacije kao i sve unutrašnje komponente na koje se oslanja. Štaviše, poruka o grešci na vrhu praćenja steka može uključivati informacije kao što su imena datoteka na strani servera i SQL kod na koji se aplikacija oslanja, omogućavajući napadaču da fino podesi injection napad.

### 2. Onesposobljena Spring CSRF zaštita
> **http.csrf().disable();**

Bez CSRF zaštite, napadači mogu da izvrše neovlašćene radnje u ime autentifikovanih korisnika.
CSRF zaštita pomaže da se osigura da zahtjevi potiču iz pouzdanih izvora, obično uključivanjem jedinstvenog tokena u svaki zahtjev koji server potvrđuje. Onemogućavanje CSRF zaštite efikasno zaobilazi ovaj sigurnosni mehanizam, omogućavajući napadačima da krivotvore zahteve bez potrebe za važećim tokenom.

### 3. Kredencijali za bazu sačuvani u okviru git sistema za praćenje verzija
> #Navode se kredencijali za konekciju na server baze
**spring.datasource.username=hak
spring.datasource.password=**
spring.h2.console.enabled=true
spring.h2.console.path=/h2-console
spring.datasource.testWhileIdle = true
spring.datasource.validationQuery = SELECT 1

U okviru sistema za kontrolu verzija, pronađen je konfiguracioni fajl koji sadrži kredencijale za pristup bazi podataka. Korištenjem ovih kredencijala, bilo ko bi mogao da pristupi podacima iz baze podataka.

### 4. Kredencijali za mail server sačuvani u okviru git sistema za praćenje verzija
> spring.mail.host=smtp.gmail.com
spring.mail.port=587
**spring.mail.username=anastasijas557@gmail.com
spring.mail.password=uitjivuciqdglsrh**
spring.mail.properties.mail.smtp.auth=true
spring.mail.properties.mail.smtp.starttls.enable=true

U okviru sistema za kontrolu verzija, pronađen je konfiguracioni fajl koji sadrži kredencijale za rad sa mail serverom. Korištenjem ovih kredencijala, bilo ko bi mogao da manipuliše slanjem i primanjem mejlova u okviru naše aplikacije. Maliciozni subjekt može da šalje lažne mejlove u ime naše aplikacije, čime je moguće navesti korisnike na oktrivanje njihovih privatnih informacija ili slati uvredljive poruke zbog kojih bi naša aplikacija izgubila korisnike ili bila krivično ugrožena.

### 5. Korisnik može da prisupi podacima drugog korisnika
U okviru aplikcije postoji autorizacija i autentifikacija, te je time ograničeno koji korisnik (po svojoj ulozi) može da pozove koju funkciju sa servera. Međutim problem je u tome što ne postoji zaštita pristupa između korisnika koji pripadaju istoj grupi. Zbog ovoga korisnik može slanjem zahtjeva našoj serverskoj aplikacji da pristupi podacima drugog korisnika.

### 6. Prenos osjetljvih informacija preko interneta
U okviru projekta korištena je komunikacija preko HTTP protokola, čime podaci koji se prenose između klijenta i servera se šalju u običnom tekstu, što ih čini ranjivim na presretanje od strane napadača (Man-in-the-Middle napadi). Ovo uključuje osjetljive informacije kao što su korisnička imena, lozinke, tokeni sesije i drugi povjerljivi podaci.

## D. Preporuke za poboljšanje koda
### 1. Otkrivanje informacija preko stack trace-a
Način da otklonimo spomenutu ranjivost je da pošaljemo korisniku opštiju poruku o grešci koja otkriva manje informacija. Takođe bismo mogli potpuno da uklonimo praćenje steka ili da logujemo poruku samo na serveru.

### 2. Onesposobljena Spring CSRF zaštita
 Kada koristimo Spring, zaštita CSRF (falsifikovanje zahtjeva na više lokacija) je podrazumjevano omogućena. Spring-ova preporuka je da koristimo CSRF zaštitu za svaki zahtjev koji bi obični korisnici mogli da obrađuju preko pretraživača.
 Ako koristimo JWT Tokene, trebalo bi da za cookie stavimo atribut SameSite na _Strict_ ili _Lax_, što može pomoći u ublažavanju CSRF napada ograničavanjem opsega cookie-a na same-site zahtjeve.
 Takođe bi bilo dobro da validiramo zaglavlja _Origin_ ili _Referer_ dolaznih zahteva da bismo bili sigurni da potiču sa očekivanog domena.

 ### 3. Kredencijali za bazu sačuvani u okviru git sistema za praćenje verzija
Aplikacija bi trebalo da skladišti svoje tajne van samog izvornog koda tj takvi podaci obično se čuvaju kroz promenljive okruženja (environment variables) ili konfiguracione datoteke. Takve datoteke ne bi trebalo da se nađu u okviru sistema za praćenje verzija jer bi postale dostupne svima koji imaju pristup tom repozitorijumu. Čuvajući ove datoteke povjerljivim, ograničava se i pristup produkcionim tajnama/kredencijalima. Ukoliko su tajne vec vidljive u okviru sistema za kontrolu verzija, neophodno je promijeniti kredencijale i ne dostaviti promjene putem sistema za kontrolu verzija.

### 4. Kredencijali za mail server sačuvani u okviru git sistema za praćenje verzija
Isto kao prethodno

### 5. Korisnik može da prisupi podacima drugog korisnika
Opisani problem može se riješiti tako što bismo vršili provjeru pristupa u okviru samih funckija. Možemo iz tokena koji se nalazi u zahtjevu da dobavimo koji korisnik šalje zahtjev i provjerimo da li mu je dozvoljeno da pristupi odgovarajućem resursu (da li su podaci vezani za njegov profil). Sam token je enkriptovan tako da nemamo problema sa Man in the Middle napadom ili izmjene informacija korisnika u okviru samoga tokena.

### 6. Prenos osjetljivih informacja preko interneta
Problem se moze riješiti korištenjem HTTPS protokola umjesto HTTP protokola. HTTPS šifruje podatke koji se prenose između klijenta i servera koristeći SSL/TLS protokole, obezbeđujući poverljivost, integritet i autentičnost.

## E. Ostale informacije

- Vrijeme provedeno pregledajući kod: 5h
- Broj identifikovanih defekata: 6


