// let postContent = document.getElementById("post-content")

// postContent.addEventListener("input", () =>{
//     autoResize(postContent)
// })

// function autoResize(textarea) {
//     textarea.style.height = 'auto'; // Reset the height to auto
//     textarea.style.height = textarea.scrollHeight + 'px'; // Set the height to match the scroll height
// }

// send request to server for adding post 

let addPostButton = document.getElementById("addPostSubmit")

addPostButton.addEventListener("click", () =>{
    let title = document.getElementById("post-title").value
    let content = document.getElementById("post-content").value
    let categories = document.querySelectorAll(".categoryInput:checked");
    let categoryValues = Array.from(categories).map((category) => category.value);
    console.log(categoryValues);                    
    let url = "/post/add/"
    let body = {
        categories: categoryValues,
        title:title,
        content:content, 
    }
    submitPost(body, url)
})

function submitPost(body, url) {
    fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(response => {
        if (response.ok) {
            window.location.href= "/"
            console.log("Succesfully created post")
        } else {
            console.error("Failed to create post");
        }
    })
    .catch(error => {
        console.error(error);
    });
}




// Check child elements on click on check box container

let checkBoxContainers = document.querySelectorAll(".checkbox-container")

checkBoxContainers.forEach(checkBoxContainer => {
    checkBoxContainer.addEventListener("click", ()=>{
        checkBox(checkBoxContainer)
    })
});

function checkBox(container) {
    let checkBoxInput = container.querySelector(".categoryInput")
    checkBoxInput.checked = !checkBoxInput.checked
}