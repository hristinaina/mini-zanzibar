## Uvod
Cilj ovoga zadatka je da konfigurisemo serversku masinu za deployment aplikacije. Aplikacija ciji deployment se vrsi je projekat implementiran za kurs IOT (Internet of Things), pod nazivom Smart Home. Ideja projekta je da imamo server ka kojem se salju podaci sa RPI uredjaja i taj server vrsi kontrolu i upravljanje nad tim podacima. Podrzan je i web interfejs u okviru aplikacije, ali za ovaj zadatak ce biti bitan samo serverski dio.

Važno je pravilno podesiti server kako bismo osigurali sigurnost, performanse i pouzdanost aplikacije. U nastavku bice opisani koraci preduzeti za konfiguraciju serverkse masine.

## Provjera operativnog sistema
Operativni sistem je bitan aspekat da bismo osigurali da koristimo najnovije verzije softvera i da su ispravke sigurnosnih propusta primjenjene. Trenutna verzija sistema na kojem ce biti vrsen deployment projekta je Ubuntu 24.04 LTS. Ova verzija je ujedno i najnovija stabilna verzija (obecava se da ce biti podrzana narednih 5 godina), sto je bitno jer znamo da ce server na kojem podizemo aplikaciju biti i u buducnosti podrzan. Ranjivosti ovoga operativnog sistema moguce je vidjeti na iducem linku: https://ubuntu.com/security/notices?order=newest&release=noble&details=

## Upravljanje vremenom
Sinhronizovanje sata sistema sa NTP-om i UTC vremenskom zonom omogucava tačno i pouzdano merenje vremena u računarskim sistemima, što je bitno za različite operativne i bezbednosne zadatke (upravljanje logovima, verifikacija SSL certifikata, autentifikacija bazirana na vemenu).
Za vremensku zonu je postavljen Belgrade/Europe komandom __sudo timedatectl set-timezone Etc/UTC__ , a nakon toga je podesen NTP server i postavljeno da se pokrece prilikom svakog pokretanja samog servera. Za to su koristene komande __sudo systemctl start ntp__ i __sudo systemctl enable ntpsec__. Za aplikaciju SmartHome je ovo od velike vaznosti jer se bavi manipulacijom podataka gdje bitan faktor predstavlja vrijeme kada su ti podaci generisani.


## Instalacija paketa
U okviru konfiguracije servera za deployment projekta bitno je instalirati neophodne pakete i provjeriti da li oni sadrze neke ranjvosti. 
Instalirani su: __TODO__

## Logging
Logovanje je konfigurisano u okviru */etc/rsyslog.conf* fajla i ostavljena je standardna konfiguracija. Time je obezbijedjeno da se logovanje vrsi samo lokalno i ne vrsi se nikakav backup na remote server.

## Konfiguracija Firewall-a:
SmartHome aplikacija (serverski dio) trci na portu 8080. Dozvoljen je pristup aplikaciji preko http komandom __sudo ufw allow http__ cime je omogucen pristup svima preko interneta. Pristup preko ssh dozvoljen je samo za odredjene ip adrese i to je izvrseno komandom __sudo ufw allow from x.x.x.x to any port 22 proto tcp__. S Obzirom da se koristi NTP, unesena je dodatna komanda za dozvolu i udp saobracaja __sudo ufw allow 123/udp__. Takodje je aktivirano logovanje komandom __sudo ufw logging on__. Ista pravila su primjenjena i na ipv6.

## Pregled file sistema


## SSH ključevi:
Kljucevi su generisani komandom __ssh-keygen -t rsa -b 4096__. Nakon sto su kljucevi geenrisani, javni kljuc je prenijet na ssh server i to se radi komandom __ssh-copy-id -i ~/.ssh/id_rsa.pub username@server_ip__. Na serveru je omogucen ssh pristup koristenjem kljuceva, a onemogucen pristup koristenjem lozinke. To se podesava u okviru */etc/ssh/sshd_config* datoteke. Nakon ovih postavki mozemo da se prijavimo na server bez unosenja lozinke, vec koristenjem naseg privatnog kljuca. Iz nekog razloga mi ssh username@server_ip sad izbacuje gresku Permission denied (publickey) :(


## Pratite performanse sistema
Logovanje da bismo pratili performanse aplikacije i znali koliko resursa treba da dodijelimo serveru.
