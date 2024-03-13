# Hešovanje lozinki

Algoritmi za heširanje imaju raznoliku primenu. Jedna od njih jeste heširanje korisničkih lozinki. Algoritam heširanja predstavlja jednosmernu kriptografsku funkciju koja generiše jedinstvenu heš vrednost za identifikaciju ulaznih podataka. Na ovaj način obezbeđuje se integritet podataka, čuva se njihova sigurnost, bez otkrivanja samih podataka. Jednom kada je nešto heširano to je praktično nepovratno jer bi bilo potrebno previše računarske snage kao i vremena.

Neki od trenutno najbezbednijih algoritama su:

**● SHA-256:** Ovaj algoritam je preporučen i odobren od strane Nacionalnog Instituta za standarde i tehnologiju (NIST). Generiše 256-bitnu heš vrednost.

**● Argon2:** Ovaj algoritam heširanja posebno dizajniran da zaštiti kredencijale. Smatra se jednim od najbezbednijih i preporučuje ga projekat bezbednosti otvorenih veb aplikacija (OWASP), pruža visok nivo odbrane od napada zasnovanih na GPU-u.

**● Bcrypt:** Ova funkcija heširanja lozinke je napravljena da uspori brute force napad. U proces heširanja uključeno je dodavanje “soli” (_password salting_), štiteći tako sačuvane heš vrednosti od napada duginih tabela (_rainbow table attack_).

**● PBKDF2** (password-based key derivation function)**:** Preporučen od strane NIST-a. Ovaj algoritam heširanja je mnogo sporiji od SHA, stoga je još jedan pogodan algoritam heširanja za zaštitu lozinki. Koristi so za generisanje heševa lozinki koje je teže pogoditi i proizvodi izlaz konfigurabilne veličine (npr. 256 bita).

Najbolji algoritam heširanja je onaj koji otežava napadačima da pronađu dve vrednosti sa istim heš izlazom. Da bismo zaštitili lozinke, stručnjaci predlažu korišćenje jakog i sporog algoritma, poput Argon2. 

Konfiguracioni parametri ovog algoritma su:

**● password** P: lozinka (ili poruka) koja će biti korišćena

**● salt** S: nasumično generisana vrednost soli (preporučeno je dužina od 16 bajtova)

**● iterations** t: broj iteracija za izračunavanje heša

**● memorySizeKB** m: količina memorije (u kilobajtima) koja će biti korišćena

**● parallelism** p: stepen paralelizma (tj. broj niti)

**● outputKeyLength** T: broj željenih vraćenih bajtova koji predstavljaju generisanu heš vrednost

Preporučena praksa korišćenja Argon2 je sa minimalnom konfiguracijom od 19 MiB memorije, brojem iteracija od 2 i stepenom paralelizma od 1.

Kada govorimo o heširanju lozinki, pouzdan provajder je onaj koji pruža sigurne i efikasne mehanizme za heširanje lozinki i sprečavanje različitih vrsta napada. Postoje različite vrste provajdera koji podržavaju Argon2 algoritam. Svaki od njih ima svoje specifičnosti i preferencije u zavisnosti od jezika ili tehnologije koju koristimo.

**● libsodium:** libsodium je popularna kriptografska biblioteka koja podržava različite moderne algoritme, uključujući Argon2. Ona je otvorenog koda, pruža visok nivo bezbednosti i lako je integrisana u različite programske jezike i platforme.
 
**● Passlib:** za rad u Python prograskom jeziku

**● Argon2.rs:** za rad u Rust programskom jeziku

**● BouncyCastle:** za rad u Java programskom jeziku

Najnovija verzija Argon2 algoritma nije identifikovana sa ozbiljnim ranjivostima koje bi
ugrozile bezbednost.

Bezbedna implementacija heš mehanizma zahteva rigorozno praćenje preporučenih praksi za konfiguraciju, odabir pouzdanog provajdera, i redovno ažuriranje na najnovije verzije algoritma kako bi se održala bezbednost sistema. Takođe, neophodno je osigurati sigurno skladištenje hešovanih lozinki i primeniti dodatne mere kao što su solenje i iterativno hešovanje radi dodatne zaštite od napada.


