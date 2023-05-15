// add comment under post
let postId = document.getElementById("id")
let addCommentButton = document.getElementById("addCommentBtn")
addCommentButton.addEventListener("click", ()=>{
    submitNewComment(postId)
})

// function editComment(id) {
    //     document.getElementById("comment-content-" + id).style.display = "none";
    //     document.getElementById("edit-comment-" + id).style.display = "block";
    // }
    
function submitNewComment(id) {
    let content = document.getElementById("new-comment").value;
    fetch("/post/comment", {
            method: "POST",
            body: JSON.stringify({ post_id: id, content: content }),
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