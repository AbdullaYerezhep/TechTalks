function editComment(id) {
    document.getElementById("comment-content-" + id).style.display = "none";
    document.getElementById("edit-comment-" + id).style.display = "block";
}
  
function submitComment(id) {
    let content = document.getElementById("edit-comment-content-" + id).value;
    fetch("/comment/edit", {
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

function deleteComment(id) {
    fetch("/comment/delete", {
        method: "POST",
        body: JSON.stringify({ id: id}),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
        location.reload()
        } else {
        console.error("Failed to delete comment");
        }
    })
    .catch(error => {
        console.error(error);
    });
}