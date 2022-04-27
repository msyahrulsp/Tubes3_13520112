import "./Button.scss";

export const Button = (props) => {
    return (
        <button
            className="button"
            onClick={(() => props.onClick(props.onClickParams))}
        >
            {props.children}
        </button>
    )
}