interface RegisterProps {
  url: string;
}

export const Register = (props: RegisterProps) => {
  return (
    <div>
      <h1>Register feeds</h1>
      <p>Register your feeds here.</p>
      <p>URL: {props.url}</p>
    </div>
  );
};
