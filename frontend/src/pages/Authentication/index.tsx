import { Container } from "@material-ui/core"
import AppBar from "@material-ui/core/AppBar"
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles"
import Toolbar from "@material-ui/core/Toolbar"
import Typography from "@material-ui/core/Typography"
import { AxiosError } from "axios"
import React, { useState } from "react"
import authentication from "../../api/authentication"
import { ApiError, ResponseStatus } from "../../api/base"
import Form, { AuthForm } from "./Form"

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
    formContainer: {
      position: "absolute",
      top: "50%",
      left: "50%",
      transform: "translate(-50%, -50%)",
    },
  })
)

function Authenticate() {
  const classes = useStyles()
  const initialFormState = { login: "", password: "" }
  const authFormState = useState<AuthForm>(initialFormState)

  const onSubmit = async ({ login, password }: AuthForm): Promise<any> => {
    try {
      const res = await authentication({ userName: login, password })
      alert("Authenticated successfully")
      console.table(res)
    } catch (e) {
      const { response } = e as AxiosError<ApiError>
      if (response?.status == ResponseStatus.UNAUTHORIZED) {
        alert("Authentication Failed")
      } else if (response?.status == ResponseStatus.BAD_REQUEST) {
        alert(response?.data.errMsg)
      }
    }
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
        <Form
          onSubmit={onSubmit}
          initialState={initialFormState}
          resetOnSubmit
          formState={authFormState}
        />
      </Container>
    </>
  )
}

export default Authenticate
