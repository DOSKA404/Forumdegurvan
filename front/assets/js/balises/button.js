export const CreateButton = (text, id, className, parent) => {
    const button = document.createElement("button");
    button.textContent = text;
    button.id = id;
    button.className = className;
    parent.appendChild(button);
    return button;
}; 