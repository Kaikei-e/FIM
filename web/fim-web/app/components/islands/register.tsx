interface RegisterProps {
  url: string;
}

export const Register = (props: RegisterProps) => {
  return (
    <div className="h-full w-full flex-col bg-white p-4 rounded-lg shadow-md">
      <div>
        <h1>Register feeds</h1>
        <p>Register your feeds here.</p>
        <p>URL: {props.url}</p>
      </div>
      <div>
        <button
          id="registerButton"
          className="bg-blue-500 text-white p-2 rounded-lg"
        >
          Register
        </button>
      </div>
    </div>
  );
};
