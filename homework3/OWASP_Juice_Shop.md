## KLASE NAPADA

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

7. **Broken Access Control**
- Klasa napada koja omogućava neovlašćenim korisnicima da pristupe, menjaju ili brišu podatke kojima ne bi trebalo da imaju pristup.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - neovlašćeno pristupanje informacijama
  - modifikacija ili uništavanje podataka
  - obavljanje funkcija van granica korisnika
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - zaobilaženje provera kontrole pristupa izmenom URL adrese
  - direktno iskorištavanje endpointa
  - pogrešna konfiguracija CORS-a 
  - nebezbedno upravljanje i manipulacija sesijama
  - neadekvatna kontrola pristupa
  - podizanje privilegija korisnika
  - Nebezbedne direktne reference objekata (IDOR)
- Primerene kontramere za sprečavanje napada:
  - implementacija principa najmanjih privilegija
  - bezbedno upravljanje sesijom
  - snažna autentifikacija i autorizacija
  - pravilno upravljanje identitetima objekata
  - logovanje grešaka u kontroli pristupa
  - brojno ograničavanje pristupa API-ju
  

8. **Security through Obscurity**
- Praksa pokušaja zaštite sistema ili podataka skrivanjem informacija o sigurnosnim mehanizmima ili dizajnu softvera. Ova praksa se oslanja na ideju da ako se informacije o sistemu drže tajnim ili su nedostupne napadačima, sistem će biti siguran.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - krađe intelektualnog vlasništva
  - krađe ili kompromitovanja osetljivih podataka korisnika
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - nedostatak transparentnosti
  - ovisnost o tajnosti
  - slabe ostale mjere zaštite
- Primerene kontramere za sprečavanje napada:
  - korištenje dodatnih sigurnosnih mehanizama
  - korištenje neobičnih portova
  - oslanjanje na manje poznate tehnike zaštite umjesto na standardne

9. **Security Misconfiguration**
  - Ova klasa odnosi se na nepravilnu konfiguraciju bezbednosnih podešavanja, dozvola i kontrola koje mogu dovesti do ranjivosti i neovlašćenog pristupa.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - neovlašćeni pristup mrežama, sistemima i podacima
  - pristup osetljivim informacijama, kao što su korisnički kredencijali, lični podaci ili poverljivi poslovni podaci
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - slaba enkripcija i heširanje
  - ostavljanje podrazumevanih korisničkih imena i lozinka nepromenjenih
  - nedostatak bezbedne komunikacije odnosno korištenje HTTP umesto HTTPS
  - loša konfiguracija http zaglavlja
  - nebezbedne dozvole za fajlove
  - prikaz poruka o grešci koje sadrže osetljive informacije
  - neadekvatne kontrole pristupa koje dozvoljavaju neovlašćenim korisnicima da vrše privilegovane radnje
  - softver je zastareo i ranjiv
  - otvoreni portovi koji nisu potrebni za poslovanje aplikacije
- Primerene kontramere za sprečavanje napada:
  - izbegavanje korišćenja podrazumevanih kredencijala
  - bezbedna komunikacija internetom odnosno adekvatna konfiguracija HTTPS-a
  - primena principa najmanjih privilegija
  - redovno ažuriranje sistema
  - bezbedno konfigurisanje http zaglavlja
  - eksportovanje samo onih aktivnosti i servisa koje su neophodne za eksportovanje

10. **XXE (XML External Entity)**
  - XXE je klasa napada koja dozvoljava napadaču da u okviru aplikacije ometa obradu XML podataka. Često dozvoljava napadaču da pregleda datoteke na serverskom sistemu datoteka i da komunicira sa bilo kojim pozadinskim ili spoljnim sistemima kojima aplikacija može da pristupi.
- Uticaj iskorištenja ranjivosti klase može dovesti do:
  - remote izvršavanje koda (pod privilegijom aplikacije)
  - pristup osetljivim podacima poput korisničkih lozinki
  - dobijanje pristupa osetljivim putanjama lokalno na serveru putem directory traversal
  - DoS, koji može dovesti do preopterećenja resursa na serveru i do uskraćivanja usluge
  - dobijanje pristupa drugim direktorijumima na mreži
  - server se može navesti da šalje HTTP zahteve na bilo koju URL adresu kojoj server može da pristupi (SSRF)
- Ranjivosti u softveru koje su dozvolile da napad uspe:
  - zastareli ili loše konfigurisani parseri
  - loša validacija i sanitizacija input-a
  - XML procesor je konfigurisan da rešava obradu spoljnih entiteta u okviru DTD (Document Type Definition)
- Primerene kontramere za sprečavanje napada:
  - primena napredne WAF (Web Application Firewall) zaštite
  - konfiguracija XML parsera da se onemogući spoljna obrada entiteta
  - korištenje sigurnih XML parsera
  - redovno ažuriranje sistema 
  - validacija ulaznih XML dokumenata

11. **Unvalidated Redirects**
  - Unvalidated Redirects je klasa napada koja dozvoljava napadaču da se redirektuje na URL-ove bez prethodne validacije i dozvole. Na ovaj način može doći do različitih napada, kao što je phishing napad.
  - Uticaj iskorištenja ove klase može dovesti do krađe korisničkih podataka, kompromitovanja sistema, širenja malicioznog softvera i drugih problema.
  - Ranjivosti u softveru koje su dovele do uspješnosti napada su nedostaci validacije i provjere URL-ova ili slaba validacija. Takođe, ranjivost može izazvati i nedostatak autentifikacije prije usmjeravanja korisnika.
  - Primjerene kontramjere:
      - validacija URL-ova
      - upotreba sigurnih preusmjeravanja
      - autentifikacija prije usmjeravanja
      - upozoravanje i informisanje korisnika o potencijalnim phishing akcijama

12. **Vulnerable components**
  - Vulnerable components klasa nastaje kada softver koristi komponente koje imaju određene sigurnosne ranjivosti. Takve komponente mogu biti zastarjele verzije biblioteka, koje napadači mogu iskoristiti za izvođenje različitih vrsta napada.
  - Uticaj iskoštenja ove klase napada može dovesti do:
      - kompromitovanja podataka
      - izvršavanje malicioznog koda
      - krađe korisničkih podataka
      - kompromitovanja sistema
  - Ranjivosti u softveru koje bi dovele do uspješnosti napada:
      - korištenje zastarjelih verzija komponenti koje imaju određene ranjivosti
      - nedostatak redovnog ažuriranja komponenti kako bi se ispravile određene ranjivosti
      -  nedostatak testiranja, praćenja i provjeravanja sigurnosti različitih komponenti sistema
  - Primjerene kontramjere:
      - redovno ažuriranje komponenti koje imaju određene ranjivosti
      - upotreba sigurnih komponenti
      - redovno testiranje sigurnosti sistema

13. **Insecure deserialization**
  - Insecure deserialization klasa napada nastaje kada softver obavlja deserijalizaciju podataka bez prethodne provjere. Napadači mogu podmetnuti serijske podatke pomoću kojih će izvršiti različite štetne akcije.
  - Uticaj iskorištenja ove klase može dovesti do:
      - krađa korisničkih podataka
      - promjena podataka
      - izvršavanje malicioznog koda
      - kompromitovanje sistema
  - Ranjivosti u softveru koje bi dovele do uspješnosti napada:
      - nedostatak provjere integriteta podataka prije deserijalizacije podataka
      - nedostatak provjere autentičnosti podataka i njihovog izvora
      - nedostatak uklanjanja podataka prije njihove deserijalizacije
  - Primjerene kontramjere:
      - validacija serijskih podataka
      - otklanjanje i čišćenje podataka prije njihove obrade
      - korištenje sigurnih biblioteka i mehanizama deserijalizacije (koje će automatski vršiti različite sigurnosne provjere)

14. **Miscellaneous**
  - Miscellaneous napad podrazumijeva napade kao što su zero-day napadi, nepravilno konfigurisani serveri, napadi od strane insider-a, loše upravljanje pristupom, nedostatak sigurnosnih ažuriranja.
  - Uticaj iskorištenja ove klase može dovesti do:
      - krađa podataka
      - neovlašćen pristup
      - izvršavanje malicionznog koda
      - kompromitovanje sistema
      - smanjena dostupnost sistema
  - Ranjivosti u softveru koje bi dovele do uspješnosti napada:
      - nedostatak ažuriranja
      - nepravilno konfigurisani serveri
      - loše obavljeno testiranje i slabe provjere sigurnosti sistema
      - nepravilno rukovanje osjetljivim podacima
      - neovlašćen pristup podacima
  - Primjerene kontramjere:
      - redovno ažuriranje
      - redovno testiranje i praćenje sigurnosti sistema
      - korištenje sigurnosnih alata, kao što su firewall, antivirusni programi...
      - edukacija korisnika o mogućim opasnostima i kako se zaštiti od istih

  15. **Broken Anti Automation**
      - Broken Anti Automation napad nastaje kada softver koristi mehanizme poput CAPTCHA-e, rate limiting-a, token-based zaštite ili druge tehnike zaštite od automatizacije, ali su ti mehanizmi neispravno implementirani ili nisu dovoljno snažni da bi zaustavili automatizovane napade.
      - Uticaj iskorištenja ove klase može dovesti do:
        - neovlašćenog pristupa
        - preopterećenja servera
        - degradacije performansi
        - kompormitovanja sistema
      - Ranjivosti u softveru koje su dovele do uspješnosti napada:
        - slabe CAPTCHA-e koje se mogu lako zaobići
        - nedostatak pravilnog rate limiting-a
        - nedostatak detekcije i sprječavanja automatizovanih napada
      - Primjerene kontramjere:
        - pravilno podešavanje CAPTCH-e
        - pravilno podešavanje rate limiting-a
        - korištenje komponenti za detekciju i sprješavanje automatizovanih napada
        - redovno ažuriranje i testiranje sigurnosti sistema

## IZAZOVI

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

Katarina:
  - Login Jim i Login Admin (Injection)
  - Kill Chatbot (Vulnerable Components)
  - Deluxe Fraud (Improper Input Validation)
  - Allowlist Bypass (Unvalidated Redirects)
  - Deprecated Interface (Security Misconfiguration)
  - Captcha Bypass (Broken Anti Automation)
