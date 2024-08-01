function getOrigin() {
    const origin = window.location.origin;
    return origin;
}

function getPredictUrl() {
    const url = `${getOrigin()}/predict`;
    return url
}

function getHomeSelect() {
    let home = document.getElementById("home-select");
    return home
}

function getGuestSelect() {
    let guest = document.getElementById("guest-select");
    return guest
}

function makePredictData(home, guest) {
    return {
        home: home.value,
        guest: guest.value,
    }
}

async function getPredict(event) {
    let home = getHomeSelect();
    let guest = getGuestSelect();
    const url = getPredictUrl();

    let data = makePredictData(home, guest);

    try {
        const response = await fetch(url, {
            method: "post",
            mode: "same-origin",
            cache: "no-cache",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        const predict = await response.json();

        let result = document.getElementById("predict-result");
        let total = document.getElementById("predict-total");

        result.textContent = `Итог: ${predict.result}`;
        total.textContent = predict.total;
    } catch (error) {
        console.log("Error:\t", error);
        alert("Что-то пошло не так...");
    }
}

function setListener() {
    let home = document.getElementById("home-select");
    let guest = document.getElementById("guest-select");
    home.addEventListener("change", getPredict);
    guest.addEventListener("change", getPredict);
}

function setTeams() {
    let teams = [
        "Германия",
        "Шотландия",
        "Венгрия",
        "Швейцария",
        "Испания",
        "Хорватия",
        "Италия",
        "Албания",
        "Словения",
        "Дания",
        "Сербия",
        "Англия",
        "Польша",
        "Нидерланды",
        "Австрия",
        "Франция",
        "Бельгия",
        "Словакия",
        "Румыния",
        "Украина",
        "Турция",
        "Грузия",
        "Португалия",
        "Чехия",
    ];

    let home = document.getElementById("home-select");
    let guest = document.getElementById("guest-select");

    for (team of teams) {
        let optionHome = document.createElement("option");
        let optionGuest = document.createElement("option");
        optionHome.text = team;
        optionHome.value = team;
        optionGuest.text = team;
        optionGuest.value = team;
        home.add(optionHome);
        guest.add(optionGuest);
    }

    setListener();
}

window.setTimeout(setTeams, 300);