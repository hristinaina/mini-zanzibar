### 1. Confidental document (1 star)
- **rješenje**: ako odemo na *About us* stranicu, vidjećemo da se tu nalazi jedan link, te klikom na njega otići ćemo na novu stranicu koja prikazuje sadržaj izvjesnog fajla. U url putanji stranice možemo primijetiti da je naveden folder *ftp*, što znači da u pretraživaču možemo kucati *http://localhost:3000/ftp* kako bismo saznali šta se još nalazi u okviru tog foldera. Između ostalih, tu se nalazi i *acquisitions.md* fajl koji predstavlja povjerljiv dokument kojeg možemo da skinemo i time riješimo ovaj izazov.
- **klasa**: Sensitive Data Exposure

### 2. Easter Egg (4 stars)
- **rješenje**: u okviru *ftp* url putanje (foldera) možemo da primijetimo da se pored ostalih fajlova  nalazi i *eastere.gg* fajl. Međutim, ne možemo ga skinuti jer nije tipa *.md* ni *.pdf*. Da bismo mogli da skinemo fajl u md formatu, u putanju dodatno unosimo *%2500.md* (vršimo Poison Null Byte). Na ovaj način uspješno možemo da skinemo easter egg fajl i riješimo ovaj izazov.
- **klasa**: Broken Access Control

### 3. Nested Easter Egg (4 stars)
- **rješenje**: U preuzetom easter egg fajlu ostavljen je string koji je enkodiran u base64 formatu, što zaključujemo s obzirom da se završava sa *==*. Kada dekodiramo taj string, dobijemo url putanju koja je šifrovana. Ako je dešifrujemo koristeći rotaciju za 13 mjesta, dobijamo iduću putanju /the/devs/are/so/funny/they/hid/an/easter/egg/within/the/easter/egg, koja je ujedno i rješenje ovog izazova.
- **klasa**: Cryptographic Issues

### 4. Poison Null Byte (4 stars)
- **rješenje**: ovaj izazov je riješen u okviru Easter Egg izazova, dodavanjem *%2500.md* na kraj url putanje. Navedni karakteri predstavljaju nulti bajt (%00 ili %2500 u URL enkodovanju) koji predstavljaju kraj stringa.
- **klasa**: Improper Input Validation


### 5. Blockchain Hype (5 stars)
- **rješenje**: istraženo je gdje se u okviru aplikacije nalazi spisak putanja, što je *main.js* fajl te je izvšen prolazak kroz svaku od putanja da bi se otkrila ona koja dovodi do sajta o Token Sales, a to je: *tokensale-ico-ea*
- **klasa**: Security through Obscurity

### 6. Client-side XSS Protection (3 stars)
- **rješenje**: U okviru login forme postoji validacija unosa nad email poljem pa nije moguće unijeti js skriptu kao vrijednost polja. Instaliran je *Burp Suite* softver koji presreće zahtjev od klijenta, te je pomoću njega umetnuta js skripta na mjesto vrijednosti email polja i takav zahtjev je proslijeđen dalje serverkoj aplikaciji. 
- **klasa**: XSS

### 7. Login Jim (3 stars)
- **rješenje**: u okviru login forme, korištena je već znana Jimova email adresa, bez poznavanja njegove šifre. Izvršen je sql injection napad tako što je na kraj njegove email adrese konkatenirano *'--* što vrsi terminaciju upita na serverskoj strani i time se ne vrši provjeravanje lozinke.
- **klasa**: Injection