let root = "http://localhost:8083"
let map;

window.onload = function () {
    window.setInterval("processRefreshDeviceNum()", 1000)
    window.setInterval("processRefreshTotalInfo()", 1000)
    window.setInterval("processRefreshAlertInfo()", 1000)

    mapboxgl.accessToken = 'pk.eyJ1IjoiY2h0eXN5cSIsImEiOiJja3FpZjZibDAxemliMm9taXlrYXVsbG5qIn0.4ATzs3-S5JC8moXOxiCNZA';
    mapboxgl.baseApiUrl = 'https://api.mapbox.com';
    map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [-122.483696, 37.833818],
        zoom: 9
    });
    map.addControl(new mapboxgl.FullscreenControl(), "top-left");
    map.addControl(new mapboxgl.ScaleControl({
        maxWidth: 80,
        unit: 'metric'
    }), "bottom-right");

    let btnDeviceInfo = document.getElementById("device-info-refresh")
    btnDeviceInfo.addEventListener("click",()=>processRefreshDeviceTable())

    let graphicInfo = document.getElementById("graphic-information")
    graphicInfo.addEventListener("click", ()=>processRefreshDeviceTable())

    let btnDeviceNameInfo = document.getElementById("device-name-info")
    btnDeviceNameInfo.addEventListener("click", ()=>processRefreshDeviceNameTable())

    let btnDeviceSearch = document.getElementById("device-search-btn")
    btnDeviceSearch.addEventListener("click",()=>processDeviceSearch())

    let btnDeviceLocate = document.getElementById("device-location-btn")
    btnDeviceLocate.addEventListener("click",()=>processDeviceLocate())

    let btnLogout = document.getElementById("btn-logout")
    btnLogout.addEventListener("click",()=>processLogout())
    
    let btnLogin = document.getElementById("btn-login")
    if(localStorage.getItem("token") === ""){
        btnLogin.style.display = "inline"
        btnLogout.style.display = "none"
    }
    else{
        btnLogin.style.display = "none"
        btnLogout.style.display = "inline"
    }

    let searchTable = document.getElementById("device-search-table")
    searchTable.style.display = "none"

    let deviceSetting = document.getElementsByClassName("device-setting")[0]
    let deviceData = document.getElementsByClassName("device-data")[0]
    let deviceLocation = document.getElementsByClassName("device-location")[0]
    if(localStorage.getItem("token") === "" || localStorage.getItem("token") == null){
        deviceSetting.style.display = "none"
        deviceData.style.display = "none"
        deviceLocation.style.display = "none"
    }
    else{
        deviceSetting.style.display = "block"
        deviceData.style.display = "block"
        deviceLocation.style.display = "block"
    }
}

function processLogout() {
    localStorage.setItem("token", "")
    location.reload()
}

function processRefreshDeviceNum () {
    fetch(root+"/v1/getOnlineDeviceNum",{
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            var deviceNumElement = document.getElementById("online-device-number")
            if(data["code"]==200){
                deviceNumElement.innerHTML = data["data"]["num"]
            }
        })
        .catch(err=>console.log(err))
}

function processRefreshTotalInfo () {
    fetch(root+"/v1/getTotalInfo",{
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            var totalInfoElement = document.getElementById("total-info")
            if(data["code"]==200){
                totalInfoElement.innerHTML = data["data"]["num"]
            }
        })
        .catch(err=>console.log(err))
}

function processRefreshAlertInfo () {
    fetch(root+"/v1/getAlertInfo",{
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            var alertInfoElement = document.getElementById("alert-info")
            var alertPerElement = document.getElementById("alert-info-percent")
            if(data["code"]==200){
                alertInfoElement.innerHTML = data["data"]["num"]
                alertPerElement.innerText = (100 * data["data"]["num"] / document.getElementById("total-info").innerText).toFixed(2) + "%"
            }
        })
        .catch(err=>console.log(err))
}

function processRefreshDeviceTable() {
    fetch(root+"/v1/getDeviceInfo",{
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            let tbody = document.getElementById("device-table").tBodies[0]
            tbody.innerHTML = null
            for(let i=0;i<data["data"].length;++i){
                let tr = document.createElement("tr")
                let th = document.createElement("th")
                th.innerText = data["data"][i]["clientId"]
                let td1 = document.createElement("td")
                td1.innerText = data["data"][i]["alertNum"]
                let td2 = document.createElement("td")
                td2.innerText = data["data"][i]["totalNum"]
                tr.appendChild(th)
                tr.appendChild(td1)
                tr.appendChild(td2)
                tbody.appendChild(tr)
            }
            var chart = Highcharts.chart('device-column-graph', {
                data: {
                    table: 'device-table'
                },
                chart: {
                    type: 'column'
                },
                title: {
                    text: 'Device Information Number Table'
                },
                yAxis: {
                    allowDecimals: false,
                    title: {
                        text: 'piece',
                        rotation: 0
                    }
                },
                tooltip: {
                    formatter: function () {
                        return '<b>' + this.series.name + '</b><br/>' +
                            this.point.y + 'piece' + this.point.name.toLowerCase();
                    }
                }
            });
        })
        .catch(err=>console.log(err))
}

function processRefreshDeviceNameTable() {
    fetch(root+"/v1/getDeviceNameInfo",{
        method: "POST",
        body: JSON.stringify({
            "token": localStorage.getItem("token")
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            if(data["code"] == 286){
                localStorage.setItem("token", "")
                alert("please login!")
                return
            }
            else if(data["code"] != 200){
                alert(data["msg"])
                return
            }
            let tbody = document.getElementById("device-table-name").tBodies[0]
            tbody.innerHTML = null
            for(let i=0;i<data["data"].length;++i){
                let tr = document.createElement("tr")
                let td1 = document.createElement("th")
                td1.innerText = data["data"][i]["clientId"]
                let td2 = document.createElement("td")
                td2.innerText = data["data"][i]["name"]
                let td3 = document.createElement("td")
                let tButton = document.createElement("button")
                tButton.innerText = "modify"
                tButton.addEventListener("click",()=>processModifyDeviceName(i, data["data"][i]["clientId"]))
                td3.appendChild(tButton)
                tr.appendChild(td1)
                tr.appendChild(td2)
                tr.appendChild(td3)
                tbody.appendChild(tr)
            }
        })
        .catch(err=>console.log(err))
}

function processModifyDeviceName(caseNum, clientId) {
    let tr = document.querySelectorAll("#device-table-name > tbody > tr")[caseNum]

    let td2 = tr.children[1];
    let tInput = document.createElement("input")
    tInput.value = td2.innerText
    td2.innerHTML = null
    td2.appendChild(tInput)

    let td3 = tr.children[2];
    let tButton = document.createElement("button")
    tButton.innerText = "save"
    tButton.addEventListener("click",()=>processSaveDeviceName(caseNum, clientId))
    td3.innerHTML = null;
    td3.appendChild(tButton);
}

function processSaveDeviceName(caseNum, clientId) {
    let tr = document.querySelectorAll("#device-table-name > tbody > tr")[caseNum]

    let td2 = tr.children[1];
    fetch(root+"/v1/modifyDeviceName",{
        method: "POST",
        body: JSON.stringify({
            "token": localStorage.getItem("token"),
            "clientId": clientId,
            "name": td2.children[0].value
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            if(data["code"] == 200){
                alert("save name success!")
                location.reload()
            }
            else if(data["code"] == 286){
                localStorage.setItem("token", "")
                alert("please login!")
                return
            }
            else{
                alert(data["msg"])
                return
            }
        })
        .catch(err=>console.log(err))
}

function processDeviceSearch() {
    let clientId = document.getElementById("device-search-id").value
    let cnt = document.getElementById("device-search-cnt").value
    fetch(root+"/v1/searchDeviceInfo",{
        method: "POST",
        body: JSON.stringify({
            "token": localStorage.getItem("token"),
            "clientId": clientId,
            "cnt": cnt
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            if(data["code"] == 286){
                localStorage.setItem("token", "")
                alert("please login!")
                return
            }
            else if(data["code"] != 200){
                alert(data["msg"])
                return
            }
            let searchTable = document.getElementById("device-search-table")
            searchTable.style.display = "block"
            let tbody = searchTable.tBodies[0]
            tbody.innerHTML = null
            for(let i=0;i<data["data"].length;++i){
                let tr = document.createElement("tr")
                let td1 = document.createElement("th")
                td1.innerText = data["data"][i]["clientId"]
                let td2 = document.createElement("td")
                td2.innerText = data["data"][i]["info"]
                let td3 = document.createElement("td")
                td3.innerText = data["data"][i]["value"]
                let td4 = document.createElement("td")
                td4.innerText = data["data"][i]["alert"]
                let td5 = document.createElement("td")
                td5.innerText = "(" + data["data"][i]["lng"].toFixed(2) + "°, " + data["data"][i]["lat"].toFixed(2) + "°)"
                let td6 = document.createElement("td")
                td6.innerText = data["data"][i]["timestamp"]
                tr.appendChild(td1)
                tr.appendChild(td2)
                tr.appendChild(td3)
                tr.appendChild(td4)
                tr.appendChild(td5)
                tr.appendChild(td6)
                tbody.appendChild(tr)
            }
        })
        .catch(err=>console.log(err))
}

let markList = []
let tLocate = null

function processDeviceLocate() {
    let clientId = document.getElementById("device-location-id").value
    let cnt = document.getElementById("device-location-cnt").value
    if(tLocate != null){
        clearInterval(tLocate)
    }
    processDeviceLocateRefresh(clientId, cnt, true)
    tLocate = setInterval(()=>processDeviceLocateRefresh(clientId, cnt, false), 2000)
}

function processDeviceLocateRefresh(clientId, cnt, isFly) {
    fetch(root+"/v1/searchDeviceInfo",{
        method: "POST",
        body: JSON.stringify({
            "token": localStorage.getItem("token"),
            "clientId": clientId,
            "cnt": cnt
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(res=>res.json())
        .then(data=>{
            if(data["code"] == 286){
                localStorage.setItem("token", "")
                alert("please login!")
                return
            }
            else if(data["code"] != 200){
                alert(data["msg"])
                return
            }
            if(isFly){
                map.flyTo({center:[
                        data["data"][0]["lng"],
                        data["data"][0]["lat"]
                    ]});
            }
            while(true){
                let mark = markList.pop()
                if(mark == null)  break;
                mark.remove()
            }
            for(let i=0;i<data["data"].length;++i){
                var mark;
                if(i === 0){
                    mark = new mapboxgl.Marker({color: 'orange'})
                }
                else if(data["data"][i]["alert"] == 1){
                    mark = new mapboxgl.Marker({color: 'red'})
                }
                else{
                    mark = new mapboxgl.Marker({color: 'green'})
                }
                mark.setLngLat([data["data"][i]["lng"], data["data"][i]["lat"]]).addTo(map)
                markList.push(mark)
            }
            let posList = []
            for(let i=0;i<data["data"].length;++i){
                posList.push([data["data"][i]["lng"], data["data"][i]["lat"]])
            }
            processMapLine(posList)
        })
        .catch(err=>console.log(err))
}

var id = 0

function processMapLine(posList) {
    if(id > 0){
        map.removeLayer('route' + (id-1).toString())
        map.removeSource('route' + (id-1).toString())
    }
    source = map.addSource('route' + id, {
        'type': 'geojson',
        'data': {
            'type': 'Feature',
            'properties': {},
            'geometry': {
                'type': 'LineString',
                'coordinates': posList
            }
        }
    })
    layer = map.addLayer({
        'id': 'route' + id,
        'type': 'line',
        'source': 'route' + id,
        'layout': {
            'line-join': 'round',
            'line-cap': 'round'
        },
        'paint': {
            'line-color': '#888',
            'line-width': 8
        }
    });
    id = id + 1
}
