const filesInput = document.getElementById("filesInput");
const folderInput = document.getElementById("folderInput");
const formatSelect = document.getElementById("format");

const filesResult = document.getElementById("filesResult");
const folderResult = document.getElementById("folderResult");

// конвертация размера
function convertSize(bytes, format) {
    switch (format) {
        case "kb":
            return (bytes / 1024).toFixed(2) + " KB";
        case "mb":
            return (bytes / 1024 / 1024).toFixed(2) + " MB";
        case "gb":
            return (bytes / 1024 / 1024 / 1024).toFixed(2) + " GB";
        default:
            return bytes + " bytes";
    }
}

async function send(files) {
    const formData = new FormData();

    for (const file of files) {
        formData.append("files", file);
    }

    const res = await fetch("/upload", {
        method: "POST",
        body: formData
    });

    return await res.json();
}

// файлы
filesInput.addEventListener("change", async (e) => {
    const data = await send(e.target.files);
    const format = formatSelect.value;

    filesResult.innerText =
        `Files: ${data.count}\nSize: ${convertSize(data.size, format)}`;
});

// папка
folderInput.addEventListener("change", async (e) => {
    const data = await send(e.target.files);
    const format = formatSelect.value;

    folderResult.innerText =
        `Files: ${data.count}\nSize: ${convertSize(data.size, format)}`;
});