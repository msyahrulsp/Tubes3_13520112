import "./Button.scss";

export const Button = (props) => {
    return (
        <button className="button">
            {props.children}
        </button>
    )
}