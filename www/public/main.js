const postBtn = document.getElementById("postImageBtn");
const searchBtn = document.getElementById("searchImagesBtn");

function renderImages(metadata) {
    const resultSection = document.getElementById("res");
    resultSection.innerHTML = "";
    metadata.forEach((value) => {
        const divElement = document.createElement("div");
        divElement.className = "image";
        divElement.id = value["Name"];
        const imgElement = document.createElement("img");
        const imgSRC = "/api/images/" + value["Name"] + "." + value["Format"];
        imgElement.setAttribute("src", imgSRC);
        const deleteBtnElement = document.createElement("button");
        deleteBtnElement.innerText = "Delete";
        deleteBtnElement.onclick = () => {
            fetch(imgSRC, {
                method: "DELETE",
            }).then(data => data.json()).then(result => {
                console.log(result);
                searchBtn.onclick()
            });
        };
        const brElement = document.createElement("br");
        divElement.appendChild(deleteBtnElement);
        divElement.appendChild(imgElement);
        divElement.appendChild(brElement);
        resultSection.appendChild(divElement);
    });
}

postBtn.onclick = () => {
    fetch("/api/images", {
        method: "POST",
        body: document.getElementById("imageInput").files[0]
    }).then(data => data.json()).then(result => console.log(result));
}

searchBtn.onclick = () => {
    let url = "/api/images/metadata?";
    const queryParams = ["minw", "maxw", "minh", "maxh", "mins", "maxs", "f"];
    queryParams.forEach((value, index) => {
        const val = document.getElementById(value)["value"];
        url += value + "=" + val;
        if (index < (queryParams.length - 1)) url += "&";
    });
    fetch(url).then(data => data.json()).then((result) => {
        renderImages(result)
    })
}
