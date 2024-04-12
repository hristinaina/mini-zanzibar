### 1. Confidental document (1 star)
- **rješenje**: kada se ode na *About us* stranicu, vidjecemo da se tu nalazi jedan link, te klikom na njega otici cemo na novu stranicu koja prikazuje sadrzaj nekog fajla. U putanji stranice mozemo primijetiti da je naveden folder *ftp*, te u pretrazivacu mozemo kucati *http://localhost:3000/ftp* kako bismo saznali sta se jos nalazi u okviru tog foldera. Izmedju ostalih, tu se nalazi i *acquisitions.md* fajl koji predstavlja povjerljiv dokument i time je ovaj izazov rijesen.
- **klasa**: Sensitive Data Exposure

### 2. Easter Egg (4 stars)
- **rješenje**: u okviru ftp foldera/putanje (kako se dolazi do njegovog sadrzaja je objasnjeno u prethodnom izazovu) mozemo da primijetimo da se izmedju ostalih, tu nalazi i *eastere.gg* fajl.Medjutim, ne mozemo ga skinuti jer nije tipa .md ni .pdf. Da bismo mogli da skinemo u md formatu, u putanju dodatno unosimo *%2500.md* (vrsimo Poison Null Byte). Na ovaj nacin uspjesno smo skinuli easter egg fajl i rijesili ovaj izazov.
- **klasa**: Broken Access Control

### 3. Nested Easter Egg (4 stars)
- **rješenje**: U preuzetom easter egg fajlu ostavljen je string koji je ocigledno enkodiran u base64 s obzirom da se zavrsava sa *==*. Kada ga dekodiramo dobijemo url putanju koja je sifrovana. Ako je desifrujemo koristeci rotaciju za 13 mjesta, dobicemo iducu putanju "/the/devs/are/so/funny/they/hid/an/easter/egg/within/the/easter/egg" koja predstavlja rjesenje ovog izazova.
- **klasa**: Cryptographic Issues

### 4. Poison Null Byte (4 stars)
- **rješenje**: ovaj izazov je rijesen u okviru Easter Egg izazova, dodavanjem *%2500.md* na kraj url putanje. Navedni karakteri predstavljaju nulti bajt (%00 ili %2500 u URL enkodovanju) koji predstavljaju kraj stringa.
- **klasa**: Improper Input Validation


### 5. Blockchain Hype (5 stars)
- **rješenje**: istrazeno je gdje se u okviru aplikacije nalazi spisak putanja, sto je *main.js* fajl te je izvrsen prolazak kroz svaku od putanja da bi se otkrila ona koja dovodi do sajta o Token Sales, a to je: *tokensale-ico-ea*
- **klasa**: Security through Obscurity

### 6. Client-side XSS Protection (3 stars)
- **rješenje**: U okviru login forme postoji validacija unosa nad email poljem pa nije moguce unijeti js skriptu kao vrijednost polja. Instaliran je Burp Suite softver koji presrece zahtjev od klijenta, pomocu njega je umetnuta js skripta na mjesto vrijednosti email polja i takav zahtjev je proslijedjen serverkoj aplikaciji. 
- **klasa**: XSS

### 7. Login Jim (3 stars)
- **rješenje**: u okviru login forme, koristena je vec znana Jimova email adresa, bez poznavanja njegove sifre. Izvrsen je sql injection napad tako sto je uz njegovu email adresu konkatenirano *'--* sto vrsi terminaciju upita na serverskoj strani i time se ni ne vrsi provjeravanje lozinke.
- **klasa**: Injection