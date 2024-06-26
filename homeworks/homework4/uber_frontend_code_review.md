# Uber (frontend)

### 1. opis projekta
Uber je mobilna i web aplikacija implemenitrana po ugledu na postojece aplikacije (Uber, Wolt, Lift, ...) koje pružaju korisnicima mogućnost da nađu prevoz i izvrše uplatu preko interneta. Ideja je da se maksimalno olakša trasnsport korisnika uz redukciju interakcije sa prevoznikom kako bi se ceo proces ubrzao, bio konzistentniji i sigurniji.

Postoje 4 tipa korisnika aplikacije: neregistorvani korisnik, registrovani korisnik, vozač i admin.

Aplikacija pruža neregistrovanim korisnicima mogućnost registracije, logina ili praćenja stanja zauzetosti svih vozila, sto se prati preko mape. Registrovani korisnici mogu da naruče vozilo/prevoz (pri čemu navode parametre poput tipa vozila, prevoza kućnih ljubimaca, bebe u autu, itd) ili da izvrše rezervaciju istog. U realnom vremenu mogu da prate kretanje vozila i da šalju poruke svom vozaču ili grupi sa kojom se voze. Nakon završetka vožnje, mogu da ostave ocjenu i komentar vozaču ili vozilu. Pored navedenih osnovnih funkcionalnosti, korisnik takođe može da vidi istoriju svih svojih akcija na sistemu, chat-ove sa ostalim korisnicima, kao i statistiku o finansijama.

Projekat se sastoji iz 3 komponente:

- backend - implementiran u programskom jeziku Java, koristenjem Spring Boot radnog okvira
- frontend - implementiran u program jeziku TypeScript, koristenjem Angular radnog okvira
- android - implementiran u programskom jeziku Java

U ovom radu davaće se revizija koda implementiranog za **frontend** komponentu sistema.


### 2. članovi razvojnog tima
- [Anastasija Savić](https://github.com/savic-a) SV 7/2020
- [Katarina Vučić](https://github.com/kaca01) SV 29/2020
- [Hristina Adamović](https://github.com/hristinaina) SV 32/2020

### 3. statička analiza koda
Nad kodom je pokrenut [Angular ESLint](https://github.com/angular-eslint/angular-eslint) alat za statičku analizu koda. Rezultat analize možete pogledati u [ovom](analiza_koda.pdf) dokumentu. 

### 4. pronađeni defekti
1. **XSS (Cross-Site Scripting)**
- Posmatrane su stranice sa različitim tipovima formi u kojima korisnici unose različite podatke, poput: **login.comoponent.ts**, **registration.component.ts**, **add-driver.component.ts**...
- Analizom istih došle se do zaključka da _input_ polja nisu validara.
- U _typescript_ kodu, nalazi se samo logika za pozivanje metoda iz servisa i eventualno provera da li su uneti svi neophodni podaci.
- Bilo ko može da unese maliciozni kod, koji bi se izvršio.
- Tokom analize i provere ispravnosti koda kao korisnički unos u polja su se unosili XSS payload-ovi, poput `<script>alert('XSS')</script>`, na osnovu čega je zaljučeno da aplikacija na neadekvatan način filtrira korisnički unos

2. **CSRF (Cross-Site Request Forgery)**
- Posmatrane su metode u različitim servisima u kojima se obrađuju HTTP zahtevi, kao i komponente pomenute u prethodnom defektu (komponente sa formama).
- Zaključeno je da ne postoji korišćenje CSRF tokena i da ni jedna forma nije adekvatno zaštićena.
 
3. **SQL Injection** 
- Ponovo su posmatrane komponente u okviru kojih su forme.
- Korisnička polja nisu adekvatno zaštićeni.
- Napadač jednostavno može da u polja unese sql upit koji će se poslati na server i izvršiti nad bazom podataka
- Na ovaj način, napadač može da dođe do vrlo osetljivih podataka, prekrši autentfikaciju i autorizaciju.

4. **Nebezbedna komunikacija** 
- Korišćen je HTTP 

5. **JWT token i sesije**
- U fajlu **auth.service.ts** impelemntirane su metode vezane za autentifikaciju i upravljanje sesijama

      localStorage.setItem("jwt", res.body.accessToken)

- Korišćenjem _localStorage-a_ sistem je osetljiviji na XSS i phishing napade.  

6. **Nebezbedna obrada podataka**
- Analizirani su svi servisi u projektu i njihove metode.
- U nekim slučajevima primećeno je da se u okviru putanje šalju osetljive informacije kao na primer korisnikov mejl
    
        resetPassword(userEmail: string resetPassword: ResetPassword): Observable<void> { 
            return this.http.put<void>(environment.apiHost + 'api/user/' + userEmail + "/resetPassword", resetPassword);
        }


### 5. preporuka poboljšanja sistema
1. **XSS (Cross-Site Scripting)**
- Implementirati validaciju i filtriranje korisničkog unosa.
- Korišćenje najnovijih verzija biblioteka i okvira koji imaju poboljšane mehanizme zaštite od XSS napada.

2. **CSRF (Cross-Site Request Forgery)**
- Implementacija i adekvatno korišćenje CSRF tokena.
- Vršenje randomizacije i ograničenje vremenskog trajanja CSRF tokena.

3. **SQL Injection** 
- Validirati podatke pre slanja na server. 
- Na ovaj način bismo izbegli sprečili da se neželjini podaci dođu do servera i izvrše akciju nad našom bazom podataka.

4. **Nebezbedna komunikacija** 
- Koristiti HTTPS umesto HTTP-a.
- Implementirati sigurnosne kriptografske mehanizme poput serrtifikata, ključeva.

5. **JWT token i sesije**
- Razmotriti čuvanje jwt tokena u HttpOnly _cookies-ima_.
- _Cookies_ sprečavaju JavaScript kod na klijentskoj strani da čitaju token.
- Osigurati da su uključeni sigurnosni HTTP headeri poput HttpOnly i Secure, kako bismo imali dodatnu zaštitu od napada.

6. **Nebezbedna obrada podataka**
- Osetljive podatke uvek slati kao deo tela zahteva, umesto kao deo URL-a.

### 6. uloženo vreme
- vreme provedeno u analizi koda: oko 3.5, 4 sata
- broj identifikovanih defekata ručnom analizom koda: 6
- broj identifikovanih defekata statičkom analizom koda: 259
