import "./Form.scss";

export const Form = (props) => {
    return (
        <div className="form-container">
            <label>{props.title}</label>
            <input 
                type={props.type}
                placeholder={props.placeholder}
                onChange={props.onChange}
            />
        </div>
    )
}