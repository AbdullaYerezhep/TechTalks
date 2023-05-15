// let postContent = document.getElementById("post-content")

// postContent.addEventListener("input", () =>{
//     autoResize(postContent)
// })

// function autoResize(textarea) {
//     textarea.style.height = 'auto'; // Reset the height to auto
//     textarea.style.height = textarea.scrollHeight + 'px'; // Set the height to match the scroll height
// }


//Check child elements on click on check box container

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