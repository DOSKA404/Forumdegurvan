export const CreateInputText = (type, placeholder, id, className, parent) => {
    const input = document.createElement('input');
    input.type = type;
    input.placeholder = placeholder;
    input.id = id;
    input.className = className;
    parent.appendChild(input);
    return input;
}