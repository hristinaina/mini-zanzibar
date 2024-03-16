# OSNOVNE INFORMACIJE
Log datoteka, u domenu softvera, je tekstualna datoteka koja bilježi događaje, aktivnosti ili informacije vezane za rad aplikacije. 
Postoje tri vrste log datoteka:
-	sistemski logovi – vode zapise o kritičnim događajima kao što su pokretanje i zaustavljanje sistema i upozorenja o kritičnim greškama
-	aplikacijski logovi – proizvod aplikacija koje bilježe svoje aktivnosti, kao što su upiti nad bazom podataka i korisničke akcije
-	sigurnosni logovi – prate sigurnosne događaje kao što su pokušaji prijavljivanja i pristup neovlašćenih korisnika

Log datoteke imaju višestruku važnost:
              -      dijagnostika – pružaju korisne informacije za dijagnostikovanje problema
-	sigurnost – zahvaljujući logovima možemo detektovati ranjivosti i druge sigurnosne probleme u sistemu
-	praćenje performansi – logovi sadrže informaciju o vremenu kada je određena akcija pokrenuta, te omogućavaju jednostavno praćenje performansi sistema (npr. vrijeme odziva)

# STRUKTURA LOG DATOTEKA
Informacije koje pružaju logovi su:
-	 KO je izvršio akciju
-	 KOJA akcija je izvršena 
-	 tačno vrijeme kada je događaj okinut

Po potrebi, mogu da sadrže i dodatne infromacije. Treba težiti konciznosti.
<br>Log datoteke imaju jasnu i konzistentnu strukuturu. Na ovaj način je omogućeno jednostavno snalaženje i pristup traženim entitetima, objekti se lakše prepoznaju i izdvajaju. Primjer konzistentne strukture je da svaki log počinje akterom, zatim opišemo akciju i na kraju loga stavimo vrijeme kada je akcija izvršena. Ovo možemo unaprijediti izdvajanjem logova različitih vrsta, npr. logove koji opisuju greške obojimo crvenom bojom. Žuta boja se koristi za upozorenja, zelena označava uspješne operacije, a plavom bojom ćemo naznačiti informativne logove.  Ukoliko naš sistem upisuje veliku količinu logova, možemo ih odvojiti u različite datoteke, po određenom kriterijumu. Struktura varira od konkretnog sistema i problema kojim se bavimo. Kako ne bi došlo do pretpranosti, trebamo težiti upisu što manjeg broja logova.
<br>Vrlo bitan aspekt ovog problema je i vremenska dužina koliko ćemo logove čuvati u memoriji. Tako na primjer, log zahtjeva možemo čuvati 6 mjeseci, ali i 3 mjeseca. Od ovoga zavisi i količina prostora koju će log datoteke zauzimati u memoriji.
<br>Osim dobre strukture log datoteke, lakšem izdvajanju događaja, doprinosi i mogućnost indeksiranja, pretrage i filtriranja. Filtriranja trebamo omogućiti po različitim kriterijumima, kao što su vremenski raspon, nivo ozbiljnosti događaja, vrsta događaja ili određeni akteri.

# VAŽNE OSOBINE
Svi logovi trebaju da ispune osobine dostupnosti, neporecivosti i integriteta. Informacije unutar logova su dostupne u svakom trenutku. Dostupnost podrazumijeva i zaštitu od gubitka podataka. Tako se osigurava da se logovi ne izgube usljed kvarova i prekida rada sistema. Neporecivost podrazumijeva da se logovi ne mogu mijenjati ili brisati nakon što su upisani. Na taj način sprječavamo manipulaciju nad podacima. Ovo se postiže pomoću metoda digitalnog potpisa ili evidencije promjena. Integritet podrazumijeva da se informacije moraju čuvati u neizmijenjenom stanju. To se postiže metodama enkripcije i kontrolom pristupa.
<br>Događaji u kojima su bitni akteri moraju biti zapisani koncizno, tako da subjekat ne može da poriče izvršenu akciju. Ovo znači da moramo jasno naznačiti ko je izvršio akciju, ali da ne otkrivamo osjetljive informacije o korisiniku. Tako npr. logovi nikada neće sadržati informacije kao što su šifra korisnika, broj kreditne kartice ili adresa na kojoj korisnik živi. Na primjer, aktera možemo navesti pomoću korisničkog imena (samo ukoliko su korisnička imena jedinstvena). 
