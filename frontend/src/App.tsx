import { Button, Container, Paper, TextField } from "@material-ui/core"
import AppBar from "@material-ui/core/AppBar"
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles"
import Toolbar from "@material-ui/core/Toolbar"
import Typography from "@material-ui/core/Typography"
import React, { ChangeEvent, FormEvent, useState } from "react"
import Send from "@material-ui/icons/Send"

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    appBar: {
      marginBottom: theme.spacing(2),
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    formCard: {
      padding: theme.spacing(4),
    },
    formContainer: {
      position: "absolute",
      top: "50%",
      left: "50%",
      transform: "translate(-50%, -50%)",
    },
    formInput: {
      margin: `${theme.spacing(1)}px 0`,
    },
  })
)

interface AuthFormState {
  login: string
  password: string
}

function formFieldHandler(
  state: AuthFormState,
  setter: Function,
  fieldName: keyof AuthFormState
) {
  return (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    let currentState: AuthFormState = { ...state }
    currentState[fieldName] = e.target.value
    setter({ ...currentState })
  }
}

function resetForm(initialFormState: AuthFormState, setter: Function) {
  setter(initialFormState)
}

function App() {
  const classes = useStyles()
  const initialFormState = { login: "", password: "" }
  const [authForm, setAuthForm] = useState(initialFormState)

  function onSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    console.log(authForm)
    resetForm(initialFormState, setAuthForm)
  }

  return (
    <>
      <AppBar className={classes.appBar} position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            Gauth
          </Typography>
        </Toolbar>
      </AppBar>
      <Container className={classes.formContainer} maxWidth="sm">
        <Paper elevation={1} className={classes.formCard}>
          <Typography variant="h5">Login</Typography>
          <form onSubmit={onSubmit}>
            <TextField
              label="User Name"
              variant="outlined"
              fullWidth
              required
              onChange={formFieldHandler(authForm, setAuthForm, "login")}
              className={classes.formInput}
              value={authForm.login}
            />
            <TextField
              type="password"
              required
              label="Password"
              variant="outlined"
              onChange={formFieldHandler(authForm, setAuthForm, "password")}
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
      </Container>
    </>
  )
}

export default App
