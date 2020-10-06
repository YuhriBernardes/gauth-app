import React, { ChangeEvent, FormEvent } from "react"
import { Button, Paper, TextField, Typography } from "@material-ui/core"
import Send from "@material-ui/icons/Send"
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles"

// import { Container } from './styles';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    formCard: {
      padding: theme.spacing(4),
    },
    formInput: {
      margin: `${theme.spacing(1)}px 0`,
    },
  })
)

export interface AuthForm {
  login: string
  password: string
}

type SetAuthFormHandler = React.Dispatch<React.SetStateAction<AuthForm>>
export type AuthFormState = [AuthForm, SetAuthFormHandler]

interface AuthFormProps {
  onSubmit(form: AuthForm): any
  formState: AuthFormState
  resetOnSubmit: boolean
  initialState: AuthForm
}

function formFieldHandler(
  [formValues, setForm]: AuthFormState,
  fieldName: keyof AuthForm
) {
  return (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    let currentState: AuthForm = { ...formValues }
    currentState[fieldName] = e.target.value
    setForm({ ...currentState })
  }
}

const Form: React.FC<AuthFormProps> = ({
  formState,
  onSubmit,
  resetOnSubmit,
  initialState,
}) => {
  const [authForm, setAuthForm] = formState
  const classes = useStyles()
  return (
    <>
      <Paper elevation={1} className={classes.formCard}>
        <Typography variant="h5">
          Login {" " + process.env.REACT_APP_SOMETHING}
        </Typography>
        <form
          onSubmit={(e: FormEvent<HTMLFormElement>) => {
            e.preventDefault()
            onSubmit(authForm)
            resetOnSubmit && setAuthForm(initialState)
          }}
        >
          <TextField
            label="User Name"
            variant="outlined"
            fullWidth
            required
            onChange={formFieldHandler(formState, "login")}
            className={classes.formInput}
            value={authForm.login}
          />
          <TextField
            type="password"
            required
            label="Password"
            variant="outlined"
            onChange={formFieldHandler(formState, "password")}
            className={classes.formInput}
            value={authForm.password}
            fullWidth
          />
          <Button
            fullWidth
            variant="contained"
            color="primary"
            endIcon={<Send />}
            className={classes.formInput}
            type="submit"
          >
            Send
          </Button>
          <Button size="small" fullWidth variant="text" color="primary">
            or register here register
          </Button>
        </form>
      </Paper>
    </>
  )
}

export default Form
