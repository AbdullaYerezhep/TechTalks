// add comment under post
let post_id = document.getElementById("post_id").innerText
let addCommentButton = document.getElementById("addCommentBtn")
let addCommentSubmitButton = document.getElementById("addCommentSubmitBtn")

let commentLikeButtons = document.querySelectorAll(".comment-likeBtn")
let commentDislikeButtons = document.querySelectorAll(".comment-dislikeBtn")
console.log(isAuthenticated);

addCommentSubmitButton.addEventListener("click", () => {
    let content = document.getElementById("new-comment").value
    let body = {content: content, post_id: +post_id}
    let url = "/comment"
    sendRequestPost(body, url)
})

addCommentButton.addEventListener("click", () => {
    signInPrompt()
})

commentLikeButtons.forEach(button => {
    let comment_id = button.getAttribute("comment-id")
    button.addEventListener("click", () => {
        if (isAuthenticated) {
            let body = { comment_id: +comment_id, islike: 1 }
            let url = "/comment/rate/"
            submitRating(body, url)
        }else{
            signInPrompt()
        }
    })
});

commentDislikeButtons.forEach(button => {
    if (isAuthenticated) {
        let comment_id = button.getAttribute("comment-id")
        button.addEventListener("click", () => {
            let body = { comment_id: +comment_id, islike: -1 }
            let url = "/comment/rate/"
            submitRating(body, url)
        })
    }else{
        signInPrompt()
    }
});

// updateComment.addEventListener("click", ()=>{
//     let body = {}
//     let url = ""
//     if (isAuthenticated){
//         signInPrompt()
//     }else{
//         sendRequest()
//     }
// })

// deleteComment.addEventListener("click", ()=> {
//     let body = {}
//     let url = ""
//     if (isAuthenticated){
//         signInPrompt()
//     }else{
//         sendRequest()
//     }
// })

// updatePost.addEventListener("click", ()=>{
//     let body = {}
//     let url = ""
//     if (isAuthenticated){
//         signInPrompt()
//     }else{
//         sendRequest()
//     }
// })

// deletePost.addEventListener("click", ()=>{
//     let body = {}
//     let url = ""
//     if (isAuthenticated){
//         signInPrompt()
//     }else{
//         sendRequest()
//     }
// })



// rateComment.addEventListener("click", ()=>{
//     let body = {}
//     let url = ""
//     if (isAuthenticated){
//         signInPrompt()
//     }else{
//         sendRequest()
//     }
// })


function sendRequestEdit(body, url) {
    let content = document.getElementById("new-comment").value;
    fetch(url, {
            method: method,
            body: JSON.stringify(body),
            headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            location.reload()
        } else {
            console.error("Failed to add comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}

function sendRequestDelete(body, url) {
    fetch(url, {
            method: "DELETE",
            body: JSON.stringify(body),
            headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            location.reload()
        } else {
            console.error("Failed to add comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}

function sendRequestPost(body, url) {
    console.log(body, url)
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
        } else {
            console.error("Failed to add comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}
