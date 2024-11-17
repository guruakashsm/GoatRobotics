const URL = 'http://localhost:8080/rpc/GOATROBOTICS';

function HideAll() {
    document.getElementById("login").style.display = 'none'
    document.getElementById("chat").style.display = 'none'
}

function OpenChat() {
    document.getElementById("chat").style.display = 'block'
    GetMessages()
    setInterval(GetMessages, 600)
}

function Login(event) {
    // Prevent form submission (page refresh)
    event.preventDefault();

    let userid = document.getElementById("userid").value;
    userid = userid.trim();  // Make sure to trim any extra spaces

    let apiUrl = URL + `/join?id=${userid}`;

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (!data.Error) {
                localStorage.setItem("userId", userid)
                HideAll()
                OpenChat()
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

document.getElementById("loginForm").addEventListener("submit", Login);

function SendMessage(event) {
    event.preventDefault(); 

    let messsage = document.getElementById("messge").value; // Corrected typo: changed "messsage" to "messge"
    messsage = messsage.trim();

    let userId = localStorage.getItem("userId");
    if (!userId) { 
        console.log("User ID is Empty So Logging Out");
        Logout();
        return;     
    }

    let apiUrl = URL + `/send?id=${userId}&message=${messsage}`;

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log(data);
            if (!data.Error) {
                document.getElementById("messge").value = "";
                GetMessages()
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

document.getElementById("messageForm").addEventListener("submit", SendMessage);

function Logout() {
    HideAll();
    document.getElementById("login").style.display = 'block';
}

function GetMessages(){

    let userId = localStorage.getItem("userId");
    if (!userId) { 
        console.log("User ID is Empty So Logging Out");
        Logout();
        return;     
    }

    let apiUrl = URL + `/messages?id=${userId}`;
    let addedRead = false

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log(data);
            if (!data.Error) {
                document.getElementById("chatMessages").innerHTML = ""
                let html = ""
                data.messages.forEach(message => {

                    if(!message.read && !addedRead && userId != message.userId){
                        addedRead = true
                        html += `
                                    <div style="display: flex;justify-content: center;padding: 10px;">
                                        <div id="unreadMessage" class="unread-message">
                                            Unread Message
                                        </div>
                                    </div>
                                `

                    }

                    if(message.userId == userId){
                        html += ` <div class="msg right-msg">
                                    <div class="msg-img" style="background-image: url('./assets/profile.png')">
                                    </div>

                                    <div class="msg-bubble">
                                        <div class="msg-info">
                                            <div class="msg-info-name">${message.userId}</div>
                                            <div class="msg-info-time">${formatDateTime(message.time)}</div>
                                        </div>

                                        <div class="msg-text">
                                            ${message.message}
                                        </div>
                                    </div>
                                </div>     
                        `
                    }else{
                        html += `
                                    <div class="msg left-msg">
                                    <div class="msg-img" style="background-image: url('./assets/profile.png')">
                                    </div>

                                    <div class="msg-bubble">
                                        <div class="msg-info">
                                            <div class="msg-info-name">${message.userId}</div>
                                            <div class="msg-info-time">${formatDateTime(message.time)}</div>
                                        </div>

                                        <div class="msg-text">
                                             ${message.message}
                                        </div>
                                    </div>
                                </div>
                               `
                    }
                    
                });
                document.getElementById("chatMessages").innerHTML = html
            }
        
        })
        .catch(error => {
            console.error('Error:', error);
        });   
}
document.getElementById("refresh").addEventListener("submit",GetMessages)



function formatDateTime(inputDateTime) {
    const date = new Date(inputDateTime);
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); 
    const year = date.getFullYear();
    let hours = date.getHours();
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const ampm = hours >= 12 ? 'PM' : 'AM';
    hours = hours % 12;
    hours = hours ? String(hours).padStart(2, '0') : '12'; 
    return `${day}.${month}.${year} ${hours}:${minutes} ${ampm}`;
}
