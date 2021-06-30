let root = "http://localhost:8083"

window.onload = function () {
    let btnLogout = document.getElementById("btn-logout")
    btnLogout.addEventListener("click",()=>processLogout())

    let btnRegister = document.getElementById("btn-submit")
    btnRegister.addEventListener("click",()=>processRegister())

    let btnLogin = document.getElementById("btn-login")
    if(localStorage.getItem("token") === ""){
        btnLogin.style.display = "inline"
        btnLogout.style.display = "none"
    }
    else{
        btnLogin.style.display = "none"
        btnLogout.style.display = "inline"
    }
}

function processLogout() {
    localStorage.setItem("token", "")
    location.reload()
}

function processRegister() {
    let username = document.getElementById("username").value
    let password = document.getElementById("password").value
    let confirmPassword = document.getElementById("confirm-password").value
    let email = document.getElementById("email").value
    if(password !== confirmPassword){
        alert("The two passwords do not match!")
        return
    }
    fetch(root+"/v1/register",{
        method: "POST",
        body: JSON.stringify({
            "username": username,
            "password": password,
            "email": email
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
        alert("register success!")
        window.location.href = "/index"
    }
    else{
        alert(data["msg"])
    }
}

