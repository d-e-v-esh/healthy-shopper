"use client";

import z from "zod";
import { useCallback } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
// @mui
import LoadingButton from "@mui/lab/LoadingButton";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
// assets
import PasswordIcon from "../../../assets/icons/PasswordIcon";
// components
import Iconify from "../../../components/Iconify";
import FormProvider from "../../../components/FormProvider";
import RHFTextField from "../../../components/RHFTextField";

// ----------------------------------------------------------------------

type FormValuesProps = {
  email: string;
};

export default function ModernForgotPasswordView() {
  const ResetPasswordSchema = z.object({
    email: z
      .string()
      .nonempty("Email is required")
      .email("Email must be a valid email address"),
  });

  const methods = useForm<FormValuesProps>({
    resolver: zodResolver(ResetPasswordSchema),
    defaultValues: {
      email: "",
    },
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const onSubmit = useCallback(async (data: FormValuesProps) => {
    try {
      await new Promise((resolve) => setTimeout(resolve, 500));
      console.info("DATA", data);
    } catch (error) {
      console.error(error);
    }
  }, []);

  const renderForm = (
    <Stack spacing={3} alignItems="center">
      <RHFTextField name="email" label="Email address" />

      <LoadingButton
        fullWidth
        size="large"
        type="submit"
        variant="contained"
        loading={isSubmitting}
        endIcon={<Iconify icon="eva:arrow-ios-forward-fill" />}
        sx={{ justifyContent: "space-between", pl: 2, pr: 1.5 }}>
        Send Request
      </LoadingButton>

      <Link
        // component={RouterLink}
        // href={paths.authDemo.modern.login}
        color="inherit"
        variant="subtitle2"
        sx={{
          alignItems: "center",
          display: "inline-flex",
        }}>
        <Iconify icon="eva:arrow-ios-back-fill" width={16} />
        Return to sign in
      </Link>
    </Stack>
  );

  const renderHead = (
    <>
      <PasswordIcon sx={{ height: 96 }} />

      <Stack spacing={1} sx={{ my: 5 }}>
        <Typography variant="h3">Forgot your password?</Typography>

        <Typography variant="body2" sx={{ color: "text.secondary" }}>
          Please enter the email address associated with your account and We
          will email you a link to reset your password.
        </Typography>
      </Stack>
    </>
  );

  return (
    <>
      <FormProvider methods={methods} onSubmit={handleSubmit(onSubmit)}>
        {renderHead}

        {renderForm}
      </FormProvider>
    </>
  );
}
