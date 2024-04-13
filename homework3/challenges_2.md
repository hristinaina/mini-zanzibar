### 1. Login Jim i Login Admin (3 stars)
- **Rješenje:** U okviru ocjena za sokove, možemo pronaći Jim-ovu email adresu. Na email adresu dodamo '-- čime zakomentarišemo ostatak koda. Na taj način bez poznavanja Jim-ove lozinke možemo pristupiti njegovom nalogu. Za admina je bilo potrebno da u polje email unesemo `' OR True --`.
- **Klasa:** Injection

### 2. Kill chatbot (5 stars)
- **Rješenje:** Ako pogledamo kod, možemo da primijetimo kako se dodaje novi korisnik. U polju username, unosom `admin"); process=null;users.addUser("1234", "user` postavljamo svoj username kao admin i dodajemo korisnika sa izmišljenim tokenom i imenom. Ovo će zbuniti chatbot-a i konstantno će vraćati isti odgovor.
- **Klasa:** Vulnerable Components

### 3. Deluxe fraud (3 stars)
- **Rješenje:** Ovo smo riješili pomoću Burp alata. Ako odemo na stranicu da dobijemo deluxe članarinu i dodamo karticu, primijetićemo dugme za plaćanje koje je disabled jer na kartici nemamo novca. Ako uđemo u inspect i izmijenimo kod tako da je dugme enabled, bićemo u stanju da pošaljemo zahtjev. Prije nego što pošaljemo zahtjev, uključićemo intercept u Burp-u. Kada uhvatimo zahtjev, izmijenimo paymentMethod iz wallet u prazan string. Slanjem ovakvog zahtjeva, dobićemo deluxe članarinu.
- **Klasa:** Improper input validation

### 4. Allowlist Bypass (4 stars)
- **Rješenje:** Za ovaj problem smo se takođe morali konsultovati sa kodom. U kodu smo naišli na funkciju isRedirectAllowed. Ta funkcija dozvoljava redirekciju ukoliko se URL u potpunosti poklapa ili ukoliko je dozvoljeni URL sadržan u poslanom URL-u. Slanjem URL-a koji sadrži u sebi jedan od dozvoljenih URL-ova, riješili smo ovaj izazov.
- **Klasa:** Unvalidated redirects

## 5. Deprecated interface (2 stars)
- **Rješenje:** Ovaj izazov je slučajno riješen. Prilikom pisanja žalbe, navodno možemo prikačiti samo PDF dokument. Međutim, ako pogledamo kod, primijetićemo da će i XML dokument dozvoljen. Kačenjem XML dokumenta, riješili smo ovaj problem.
- **Klasa:** Security misconfiguration

## 6. Captcha Bypass (3 stars)
- **Rješenje:** Uđemo na stranicu gdje se daje review. Prije nego što potvrdimo review, uključimo intercept na Burp-u. Nakon što zahtjev stigne, poslaćemo ga repeater-u. Tu ćemo namjestiti da se zahjev ponovi 10 ili više puta. Ovo je takođe moguće riješiti i ručno, samo je bitno poslati više od 9 zahtjeva za manje od 20 sekundi.
- **Klasa:** Broken Anti Automation
