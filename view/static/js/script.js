function editComment(id) {
    document.getElementById("comment-content-" + id).style.display = "none";
    document.getElementById("edit-comment-" + id).style.display = "block";
}

let userID = document.getElementById("user_id").getAttribute("user-id");
const isAuthenticated = userID !== null && userID !== "";

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
            let url = "/post/rate/"
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
            let url = "/post/rate/"
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

