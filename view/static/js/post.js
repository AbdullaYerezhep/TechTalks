// add comment under post
let postID = document.getElementById("id").innerText
let userID = document.getElementById("user_id").innerText
console.log(postID, userID);
let addCommentButton = document.getElementById("addCommentBtn")
addCommentButton.addEventListener("click", ()=>{
    if (userID.length == 0){
        signInPrompt()
    }else{
        submitNewComment(postID)
    }
})

function signInPrompt() {
    signInContainer.classList.remove("hidden");
    overlay.style.display = "block"
}


function submitNewComment(p_id) {
    let content = document.getElementById("new-comment").value;
    fetch("/post/comment", {
            method: "POST",
            body: JSON.stringify({ post_id: p_id, content: content }),
            headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            location.reload()
        } else {
            console.error("Failed to update comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}
  
function submitUpdatedComment(id) {
    // let loc = location.href
    var content = document.getElementById("edit-comment-content-" + id).value;
    fetch("/post/comment/edit/", {
        method: "POST",
        body: JSON.stringify({ id: id, content: content }),
        headers: {
        "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
        location.reload()
        } else {
        console.error("Failed to update comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}