let root = "http://localhost:8083"

window.onload = function () {
    let btnLogin = document.getElementById("btn-login")
    btnLogin.addEventListener("click",()=>processLogin())
}

function processLogin() {
    let usernameInput = document.getElementById("username")
    let passwordInput = document.getElementById("password")
    fetch(root+"/v1/login",{
        method: "POST",
        body: JSON.stringify({
            "username": usernameInput.value,
            "password": passwordInput.value
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>processData(data))
        .catch(err=>console.log(err))
}

function processData(data) {
    if(data["code"] == 200){
        localStorage.setItem("token", data["token"])
        window.location.href = "/index"
    }
}
