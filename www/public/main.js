const postBtn = document.getElementById("postImageBtn");
const searchBtn = document.getElementById("searchImagesBtn");

const resultSection = document.getElementById("res");
const postMsg = document.getElementById("postMsg");
const resMsg = document.getElementById("resMsg");

const imageInput = document.getElementById("imageInput");

imageInput.onchange = () => {
    postMsg.innerText = "";
};

function renderImages(metadata) {
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
                searchBtn.onclick()
            });
        };
        divElement.appendChild(document.createElement("br"));
        divElement.appendChild(deleteBtnElement);
        divElement.appendChild(document.createElement("br"));
        divElement.appendChild(imgElement);
        divElement.appendChild(document.createElement("br"));
        resultSection.appendChild(divElement);
        divElement.appendChild(document.createElement("br"));
    });
}

postBtn.onclick = () => {
    postMsg.removeAttribute("class");
    postMsg.innerText = "Adding image...";
    fetch("/api/images", {
        method: "POST",
        body: imageInput["files"][0]
    }).then((result) => {
        if (result.status === 200) {
            postMsg.removeAttribute("class");
            postMsg.innerText = "Image added";
        } else {
            postMsg.setAttribute("class", "error");
            postMsg.innerText = "Unable to add image";
        }
    }).catch((_) => {
        postMsg.setAttribute("class", "error");
        postMsg.innerText = "Unable to add image";
    });
};

searchBtn.onclick = () => {
    let url = "/api/images/metadata?";
    const queryParams = ["minw", "maxw", "minh", "maxh", "mins", "maxs", "f"];
    queryParams.forEach((value, index) => {
        const val = document.getElementById(value)["value"];
        url += value + "=" + val;
        if (index < (queryParams.length - 1)) url += "&";
    });
    fetch(url).then(result => {
        if (result.ok) return result.json();
        else {
            resMsg.setAttribute("class", "error");
            resMsg.innerText = "Unable to find images";
        }
    }).then((data) => {
        resMsg.removeAttribute("class");
        resMsg.innerText = "";
        renderImages(data);
    }).catch((_) => {
        resMsg.setAttribute("class", "error");
        resMsg.innerText = "Unable to find images";
    });
};
