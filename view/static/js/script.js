function editComment(id) {
    document.getElementById("comment-content-" + id).style.display = "none";
    document.getElementById("edit-comment-" + id).style.display = "block";
}

postLikeButtons = document.querySelectorAll(".post-likeBtn")
postDislikeButtons = document.querySelectorAll(".post-dislikeBtn")

commentLikeButtons = document.querySelectorAll(".comment-likeBtn")
commentDislikeButtons = document.querySelectorAll(".comment-dislikeBtn")

console.log(commentLikeButtons);
postLikeButtons.forEach(button => {
    let post_id = button.getAttribute("post-id")
    button.addEventListener("click",() =>{
        let body = {post_id:+post_id, islike:1}
        let url = "/post/rate/"
        console.log(body, url);
        submitRating(body, url, 1)
    })
});

postDislikeButtons.forEach(button => {
    let post_id = button.getAttribute("post-id")
    button.addEventListener("click", () => {
        let body = { post_id: +post_id, islike: -1 }
        let url = "/post/rate/"
        submitRating(body, url)
    })
});

commentLikeButtons.forEach(button => {
    let comment_id = button.getAttribute("comment-id")
    button.addEventListener("click", () => {
        let body = { comment_id: +comment_id, islike: 1 }
        let url = "/comment/rate/"
        submitRating(body, url)
    })
});

commentDislikeButtons.forEach(button => {
    let comment_id = button.getAttribute("comment-id")
    button.addEventListener("click", () => {
        let body = { comment_id: +comment_id, islike: -1 }
        let url = "/comment/rate/"
        submitRating(body, url)
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

