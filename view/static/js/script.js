function editComment(id) {
    document.getElementById("comment-content-" + id).style.display = "none";
    document.getElementById("edit-comment-" + id).style.display = "block";
}

let userID = document.getElementById("user_id").getAttribute("user-id");
const isAuthenticated = userID !== null && userID !== "" && userID !== "0";

let postLikeButtons = document.querySelectorAll(".post-likeBtn")
let postDislikeButtons = document.querySelectorAll(".post-dislikeBtn")

function signInPrompt() {
    signInContainer.classList.remove("hidden");
    overlay.style.display = "block"
}

postLikeButtons.forEach(button => {
    let post_id = button.getAttribute("post-id")
    button.addEventListener("click",(event) => {
        event.stopPropagation();
        if (isAuthenticated) {
            let body = {post_id:+post_id, islike:1}
            let url = "/post/rate"
            console.log(body, url);
            submitRating(body, url, 1)
        }else{
            signInPrompt()
        }
    })
});

postDislikeButtons.forEach(button => {
    let post_id = button.getAttribute("post-id")
    button.addEventListener("click", (event) => {
        event.stopPropagation();
        if (isAuthenticated) {
            let body = { post_id: +post_id, islike: -1 }
            let url = "/post/rate"
            submitRating(body, url)
        }else{
            signInPrompt()
        }
    })
});


function submitRating(body, url) {
    fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            location.reload()
            console.log("Succesfully rated comment")
        } else {
            console.error("Failed to update comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}

const headerSignOut = document.querySelector(".headerSignOut")

const menuSignOut = document.querySelector(".menuSignOut")


if (isAuthenticated){
    headerSignOut.addEventListener("click", ()=>{
        sendPostRequest()
    })
    
    menuSignOut.addEventListener("click", ()=>{
        sendPostRequest()
    })
}



function sendPostRequest() {
    var url = "/logout";
  
    fetch(url, {
      method: "POST",
    })
    .then(response => {
        window.location.href = "/"
    })
    .catch(error => {
      console.error(error);
    });
}

function isEmpty(content) {
    if (content === "" || content === null || content === undefined) {
        return true
    }
    return false
}