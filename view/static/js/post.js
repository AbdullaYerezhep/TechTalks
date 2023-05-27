// GLOBALS

     const post_id = document.getElementById("post_id").innerText
     const post_user_id = document.getElementById("post_user_id").innerText

     const addCommentButton = document.getElementById("addCommentBtn")
     const addCommentSubmitButton = document.getElementById("addCommentSubmitBtn")

     const commentLikeButtons = document.querySelectorAll(".comment-likeBtn")
     const commentDislikeButtons = document.querySelectorAll(".comment-dislikeBtn")

     const commentAuthorElements = document.querySelectorAll(".comment-author");
     const commentAuthorIDs = Array.from(commentAuthorElements).map(element => element.innerText);

     const isMyPost = userID === post_user_id
     const hasMyComments = commentAuthorIDs.some(comment => comment.includes(userID))
     
     const postContent = document.querySelector(".single-post-content")
     const editPostContent = document.getElementById("edit-single-content")
     console.log(isAuthenticated);
     


    //POSTS
    //check if opened post author is equal to signed user
if (isMyPost) {
    const editPostButton = document.getElementById("update-post")
    const deletePostButton = document.getElementById("delete-post")

    const updatePostSubmit = document.getElementById("submit-update-post")
    const discardPostSubmit = document.getElementById("discard-update-post")

    const editContainer = document.getElementById("edit-container")
    const decisionContainer = document.getElementById("decision-container")

    //function to open edit fields
    editPostButton.addEventListener("click", () => {
        editContainer.classList.add("hidden")
        decisionContainer.classList.remove("hidden")

        postContent.classList.add("hidden")
        editPostContent.classList.remove("hidden")
    })

    //function to submit edited post    
    updatePostSubmit.addEventListener("click", ()=> {

        editContainer.classList.remove("hidden")
        decisionContainer.classList.add("hidden")

        let title = document.getElementById("edit-post-title").value
        let content = document.getElementById("edit-post-content").value

        let body = {
            id:         +post_id,
            title:      title,
            content:    content,
        }

        let url = "/post/edit"

        sendRequestEdit(body, url)

        postContent.classList.remove("hidden")
        editPostContent.classList.add("hidden")
    })


    //function to cancel and discard changes
    discardPostSubmit.addEventListener("click", ()=>{
        let title = document.getElementById("edit-post-title")
        let content = document.getElementById("edit-post-content")

        title.value = ""
        content.value = ""
        editContainer.classList.remove("hidden")
        decisionContainer.classList.add("hidden")

        postContent.classList.remove("hidden")
        editPostContent.classList.add("hidden")
    })

    //function to delete post on click
    deletePostButton.addEventListener("click", ()=>{
        let body = {
            id: +post_id
        }
        let url = "/post/delete"
        sendRequestDelete(body, url, "/")
    })
}


//COMMENTS

//function to add comment on click 
addCommentSubmitButton.addEventListener("click", () => {
    let content = document.getElementById("new-comment").value
    let body = {
        content: content, post_id: +post_id
    }
    let url = "/comment"
    document.getElementById("new-comment").value = '';
    sendRequestPost(body, url)
})


//function to prompt sign in on click 
addCommentButton.addEventListener("click", () => {
    signInPrompt()
})


//function to like comments on click
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



//function to dislike comments on click
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



// if post has comments which author is equal to signed user 
if (hasMyComments) {
    const editCommentButtons   = document.querySelectorAll(".edit-comment");
    const deleteCommentButtons = document.querySelectorAll(".delete-comment");
    const updateCommentInputs    = document.querySelectorAll(".update-comment-input");
    const editCommentInputs    = document.querySelectorAll(".edit-comment-input");
    const submitButtons        = document.querySelectorAll(".submit-update-comment");
    const discardButtons       = document.querySelectorAll(".discard-update-comment");
    const commentContents      = document.querySelectorAll(".my-comment-content");
    const commentDetails       = document.querySelectorAll(".my-comment-details");
    const commentIDsElements= document.querySelectorAll(".comment-id")
    const commentIDs = Array.from(commentIDsElements).map(element => element.innerText);


    editCommentButtons.forEach((button, index) => {
        const details = commentDetails[index];
        const updateInput = updateCommentInputs[index];
        const editInput = editCommentInputs[index];
        const submitButton = submitButtons[index];
        const discardButton = discardButtons[index];
        const commentContent = commentContents[index].innerText;
        const mycomment = commentContents[index]
        const commentID = commentIDs[index]

        button.addEventListener("click", () => {
            details.open = false;
            editInput.classList.remove("hidden");
            mycomment.classList.add("hidden")
        });

        submitButton.addEventListener("click", () => {
            const updatedComment = updateInput.value;
            editInput.classList.add("hidden");
            mycomment.classList.remove("hidden")
            let body = {
                id:         +commentID,
                content:    updatedComment
            }

            let url = "/comment/edit"
            sendRequestEdit(body, url)
        });

        discardButton.addEventListener("click", () => {
            updateInput.value = commentContent;
            editInput.classList.add("hidden");
            mycomment.classList.remove("hidden")
        });
    
        
    });
    
    deleteCommentButtons.forEach((deleteButton, index) => {
        const details = commentDetails[index];
        const commentID = commentIDs[index]
        deleteButton.addEventListener("click", () => {
            details.open = false;
            const body = {
                id:         +commentID,
                post_id:    +post_id,
            }
            const url = "/comment/delete"
            sendRequestDelete(body, url)

        });
    });
    
}


//REQUEST TO THE SERVER

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
            alert("Succesfully updated")
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

function sendRequestDelete(body, url, endpoint) {
    fetch(url, {
            method: "DELETE",
            body: JSON.stringify(body),
            headers: {
            "Content-Type": "application/json"
            }
    })
    .then(response => {
        if (response.ok) {
            if (endpoint === undefined) {
                location.reload()
            }else{
                window.location.href = endpoint
            }
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
