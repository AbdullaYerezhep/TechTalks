// add comment under post
let post_id = document.getElementById("post_id").innerText
let post_user_id = document.getElementById("post_user_id").innerText

let addCommentButton = document.getElementById("addCommentBtn")
let addCommentSubmitButton = document.getElementById("addCommentSubmitBtn")

let commentLikeButtons = document.querySelectorAll(".comment-likeBtn")
let commentDislikeButtons = document.querySelectorAll(".comment-dislikeBtn")

let postContent = document.querySelector(".single-post-content")
let editPostContent = document.getElementById("edit-single-content")

const isMyPost = userID === post_user_id
if (isMyPost){
    const editPost = document.getElementById("update-post")
    const deletePost = document.getElementById("delete-post")

    const updatePostSubmit = document.getElementById("submit-update-post")
    const discardPostSubmit = document.getElementById("discard-update-post")

    const editContainer = document.getElementById("edit-container")
    const decisionContainer = document.getElementById("decision-container")

    editPost.addEventListener("click", () => {
        editContainer.classList.add("hidden")
        decisionContainer.classList.remove("hidden")

        postContent.classList.add("hidden")
        editPostContent.classList.remove("hidden")
    })

    updatePostSubmit.addEventListener("click", ()=>{
        editContainer.classList.remove("hidden")
        decisionContainer.classList.add("hidden")

        let title = document.getElementById("edit-post-title").value
        let content = document.getElementById("edit-post-content").value
        let body = {
            id: +post_id,
            title: title,
            content: content,
        }

        let url = "/post/edit"

        sendRequestEdit(body, url)

        postContent.classList.remove("hidden")
        editPostContent.classList.add("hidden")
    })

    discardPostSubmit.addEventListener("click", ()=>{
        editContainer.classList.remove("hidden")
        decisionContainer.classList.add("hidden")

        postContent.classList.remove("hidden")
        editPostContent.classList.add("hidden")
    })

    deletePost.addEventListener("click", ()=>{
        let body = {id: post_id}
        let url = "/post/delete"
        sendRequestDelete(body, url)
    })

}

addCommentSubmitButton.addEventListener("click", () => {
    let content = document.getElementById("new-comment").value
    let body = {content: content, post_id: +post_id}
    let url = "/comment"
    content.value = ''
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
            let url = "/comment/rate"
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
            let url = "/comment/rate"
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
    fetch(url, {
            method: "PATCH",
            body: JSON.stringify(body),
            headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            location.reload()
        } else {
            alert("Failed to update")
            console.error("Failed to update");
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
            console.error("Failed to delete");
            alert("failed to delete")
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
