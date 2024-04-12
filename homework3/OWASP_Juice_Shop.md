OWASP Juice Shop

**KLASE NAPADA**

1. **Broken Authentication**
- Klasa napada koja se odnosi na ranjivosti u mehanizmima autentifikacije, kao što su lozinke, tokeni sesija, *cookies*... Ovaj tip napada omogućava napadačima da zaobiđu autentifikacione mehanizme i steknu neovlašćen pristup korisničkim nalozima.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - neovlašćenog pristupa korisničkim nalozima sa privilegijama, kao što je administratorov nalog
  - krađe ili kompromitovanja osetljivih podataka korisnika
  - mogućnosti manipulacije ili brisanja podataka korisnika
  - povrede privatnosti korisnika i poverljivih informacija
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - loša implementacija autentifikacionog sistema
  - nedovoljno jake lozinke
  - nedostatak dvofaktorske autentifikacije
  - neadekvatna zaštita administratorovog naloga
- Primerene kontramere za sprečavanje napada:
  - korišćenje snažnih lozinki za svaki korisnički nalog
  - implementacija dvofaktorske autentifikacije za dodatni sloj sigurnosti
  - redovno ažuriranje softvera i primena ispravki koje popravljaju poznate ranjivosti
  - implementacija različitih autentifikacionih mehanizama, kao što su zaključavanje naloga nakon nekoliko neuspelih pokušaja prijave
  - redovno proveravanje i nadgledanje autentifikacionih aktivnosti kako bi se otkrile nepravilnosti ili sumnjive radnje.
 
    
2. **Injection**
-  Ova klasa napada predstavlja situacije u kojoj napadač ubacuje zlonamerni ili nevalidni ulaz u aplikaciju kako bi iskoristio ranjivost i izvršio neželjene akcije. U slučaju sql injection-a, napadač ubacuje sql kod u input polje ili url parametre kako bi manipulisao sql upitima i izvršio neautorizovane akcije nad bazom podataka.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - neovlašćenog pristupa nad podacima iz baze podataka
  - krađe osetljivih informacija kao što su korisnička imena i lozinke
  - mogućnosti manipulacije ili brisanja podataka korisnika
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - neadekvatna validacija input polja, to jest korisničkog unosa
  - neadekvatno čišćenje i filtriranje unesenih podatak, pre njihovog korišćenja u aplikaciji
  - korišćenje dinamičkih, a ne parametrizovanih sql upita
- Primerene kontramere za sprečavanje napada:
  - korišćenje parametrizovanih sql upita umesto dinamičkih
  - validacija i filtriranje korisničkog unosa kako bi se uklonile potencijalno opasne komande
  - korišćenje ORM (*Object-Relational Mapping*) biblioteka koje automatski tretiraju input parametre kao podatke, a ne kao deo SQL upita.
  - implementacija principa najmanjih privilegija (*Least Privilege Principle*)
 
    
3. **Improper Input Validation**
- Klasa se odnosi na situacije kada softver neadekvatno validira ulazne podatke koje prima od korisnika ili drugih izvora. To može dovesti do različitih bezbednosnih ranjivosti, uključujući SQL Injection, XSS (Cross-Site Scripting), Command Injection...
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - neovlašćenog pristupa podacima
  - otkrivanje osetljivih informacija
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - neadekvatna validacija ulaznih podataka koji se koriste u manipulaciji sa fajlovima ili putanjama do fajlova
  - neadekvatno tretiranje null bajt (\0) u korisničkom unosu
- Primerene kontramere za sprečavanje napada:
  - validacija ulaznih podataka
  - upotreba sigurnih funkcija za rad sa fajlovima i provera dozvola pristupa
 
    
4. **XSS (*Cross-Site Scripting*)**
- Klasa obuhvata napade na veb aplikacije, gde napadač ubacuje zlonamerni JavaScript kod. Kod se izvršava na strani korisnika kada posećuje zaraženu veb stranicu.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - izvršavanja zlonamernih akcija u ime korisnika
  - krađe osetljivih podataka
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - neadekvatna validacija korisničkog unosa
  - nedostatak pravilne konfiguracije *Content Security Policy* (CSP) politika
  - slaba zaštita od CSRF (*Cross-Site Request Forgery*) napada
  - slaba kontrola pristupa i autentifikacija
- Primerene kontramere za sprečavanje napada:
  - implementacija adekvatne validacije korisničkog unosa na serverskoj strani
  - korišćenje CSP politika za ograničavanje izvršavanja skripti i drugih resursa na stranici
  - korišćenje mehanizama zaštite od CSRF napada, kao što su CSRF tokeni
  - korišćenje najnovijih verzija biblioteka i okvira koji imaju poboljšane mehanizme zaštite od XSS napada
 
    
5. **Cryptographic Issues**
- Klasa napada koja se odnosi na situacije kada napadač pokušava da manipuliše ili falsifikuje kriptografske mehanizme. Na taj način želi da dođe do neovlašćenog pristupa i izvrši neke neželjene radnje. Klasa obuhvata sve vrste napada koji su povezani sa kriptografskim procesima ili implementacijama u softveru. Na primer manipulacija digitalnih potpisa, napad na algoritme enkripcije ili dekripcije...
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - manipulacije podataka unutar sistema
  - krađe osetljivih podataka
  - izvršavanja neovlašćenih radnji
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - slabi kriptografski algoritmi
  - loše upravljanje ključevima i sertifikatima
  - nedostatak provere autentičnosti i integriteta podataka
- Primerene kontramere za sprečavanje napada:
  - implementacija sigurnih kriptografskih algoritama i protokola
  - pravilno upravljanje ključevima i sertifikatima, uključujući periodično rotiranje ključeva
  - provera autentičnosti i integriteta podataka korišćenjem digitalnih potpisa ili MAC (*Message Authentication Code*)
 
    
6. **Sensitive Data Exposure**
- Klasa napada koja se odnosi na situacije kada osetljivi podaci (lozinke, finansijski podaci ili lične informacije) nisu adekvatno zaštićeni i postaju dostupni neovlašćenim entitetima.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - narušavanja privatnosti korisnika
  - krađe identiteta
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - nedovoljno šifrovanje osetljivih podataka
  - nedostatak mehanizama za zaštitu podataka
  - nebezbedno upravljanje sesijama
  - neadekvatna kontrola pristupa
  - loše konfigurisane postavke pristupa
- Primerene kontramere za sprečavanje napada:
  - korišćenje jakog šifrovanja prilikom čuvanja osetljivih podataka
  - implementacija sigurnih protokola za prenos osetljivih podataka preko mreže
  - pravilno upravljanje identitetima i pristupom, uključujući snažnu autentikaciju i autorizaciju
  - implementacija sistema za detekciju i sprečavanje neovlašćenog pristupa osetljivim podacima, poput IPS/IDS sistema ili SIEM platformi


**IZAZOVI**

Anastasija: 
  - Password Strength (Broken Authentication)
  - User Credentils (Injection)
  - Poison Null Byte (Improper Input Validation)
  - Server-side XSS Protection (XSS)
  - Forged Coupon (Cryptographic Issues)
  - Misplaced Signature File (Sensitive Data Exposure)

Hristina: 
  - Confidental document (Sensitive Data Exposure)
  - Easter Egg (Broken Access Control)
  - Nested Easter Egg (Cryptographic Issues)
  - Poison Null Byte (Improper Input Validation)
  - Blockchain Hype (Security through Obscurity)
  - Client-side XSS Protection (XSS)
  - Login Jim (Injection)     