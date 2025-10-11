import { Component, effect, ElementRef, input, output, untracked, viewChild } from "@angular/core";
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

import { take, tap } from "rxjs";

import { AutoFocusModule } from 'primeng/autofocus';
import { ButtonModule } from "primeng/button";
import { DialogModule } from 'primeng/dialog';
import { TextareaModule } from 'primeng/textarea';

import { UpdateUserBioRequestBody, UserMeService } from "@cubeshares/services/user";

type UpdateUserBioForm = {
  [K in keyof UpdateUserBioRequestBody]: FormControl<UpdateUserBioRequestBody[K]>;
};

@Component({
  host: {
    '(keydown.meta.enter)': 'onSubmit()',
    '(keydown.control.enter)': 'onSubmit()',
  },
  selector: 'cubeshares-user-bio-edit-dialog',
  templateUrl: 'user-bio-edit-dialog.component.html',
  styleUrl: 'user-bio-edit-dialog.component.scss',
  imports: [AutoFocusModule, ButtonModule, DialogModule, ReactiveFormsModule, TextareaModule],
})
export class UserBioEditDialogComponent {
  readonly currentValue = input.required<string>();
  readonly visible = input<boolean>(false);
  readonly visibleChange = output<boolean>();

  readonly bioTextarea = viewChild<ElementRef<HTMLTextAreaElement>>('bioTextarea');

  protected form: FormGroup<UpdateUserBioForm> | undefined;

  constructor(private readonly fb: FormBuilder, private readonly userMeService: UserMeService) {
    effect(() => {
      const value = this.currentValue()
      untracked(() => {
        if (value === undefined) return;
        this.form = this.fb.group<UpdateUserBioForm>({
          bio: this.fb.control(this.currentValue(), { nonNullable: true }),
        });
      })
    })
  }

  protected onSubmit(): void {
    if (!this.form) return;
    const requestBody = this.form.getRawValue();
    requestBody.bio = requestBody.bio.trim();
    this.userMeService.updateUserBio(requestBody).pipe(
      take(1),
      tap(response => {
        if (response !== null) return;
        this.userMeService.pollReadUserMe();
        this.form?.reset();
        this.visibleChange.emit(false);
      })
    ).subscribe();
  }

  protected onShowFocus(): void {
    this.bioTextarea()?.nativeElement.focus();
  }
}
