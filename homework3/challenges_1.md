### 1. Password Strength (2 star)
- **rešenje**: na main stranici nadjemo review koji je ostavio admin i tako pronađemo njegom email. Lozinka je brute force-ovana (bilo je lako sa obzirom da se nije potrudio da lozinka bude komplikovana :D) (email: admin@juice-sh.op lozinka: admin123)
- **klasa**: Broken Authentication

### 2. User Credentials (4 stars)
- **rešenje**: (http://127.0.0.1:3000/rest/products/search/), putanja predstavlja link do spiska svih proizvoda u jiuce shop-u. Pretpostavljamo da u bazi postoji tabela users i pokušavamo da izvršimo sql injection. upit: ?q=')) union select id,email,password,4,5,6,7,8,9 from users--
- **klasa**: Injection

### 3. Poison Null Byte (4 stars)
- **rešenje**: ovaj izazov je rešen dodavanjem *%2500.md* na kraj url putanje. Navedni karakteri predstavljaju nulti bajt (%00 ili %2500 u URL enkodovanju) koji predstavljaju kraj stringa.
- **klasa**: Improper Input Validation

### 4. Server-side XSS Protection (4 stars)
- **rešenje**: U tekstu zadatka je postavljen deo koda koji treba da se umetne. Na kraju u polje za komentare ostavljamo naš maliciozni kod i svaki put kada se otvori starnica za prikaz komentara, izvršiće se naš kod.
- **klasa**: XSS

### 5. Forged Cupon (6 stars)
- **rešenje**: u okviru *ftp* url putanje, pomoću Poison Null Byte, skidamo fajlove coupons_2013.md. U tom fajlu pronalazimo listu enkodovanih kupona. Preuzimamo i fajl package.json, iz kog u delu dependency, pronalazimo biblioteku koja se koristi za enkodovanje podataka. Vidimo da se koristi z85 i dekodujemo jedan od kupona kako bismo saznali format kupona i kreiramo svoj lažni :D. Format je MESECGOD-VREDNOSTKUPONA (APR24-90)
- **klasa**: Cryptographic Issues

### 6. Misplaced Signature File (4 stars)
- **rešenje**: u okviru *ftp* url putanje, pomoću Poison Null Byte, skidamo fajl suspicious_errors.yml
- **klasa**: Sensitive Data Exposure