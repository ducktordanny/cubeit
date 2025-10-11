import { effect, inject, Injectable, signal, untracked } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';

import { MessageService } from 'primeng/api';
import {
  catchError,
  of,
  Observable,
  shareReplay,
  switchMap,
  take,
  tap,
  timer,
  finalize,
  takeUntil,
  Subject,
} from 'rxjs';

import { ApiService } from '../api';
import { UpdateUserBioRequestBody, UserResponse } from './user.type';

@Injectable({ providedIn: 'root' })
export class UserMeService {
  readonly loggedInUser = signal<UserResponse | null>(null);
  readonly isLoading = signal<boolean>(true);
  readonly error = signal<HttpErrorResponse | null>(null);

  private readonly resetPreviousPoll = new Subject<void>();

  private readonly api = inject(ApiService);
  private readonly router = inject(Router);
  private readonly messageService = inject(MessageService);

  constructor() {
    this.pollReadUserMe();
    this.observeErrorState();
  }

  pollReadUserMe(): void {
    this.resetPreviousPoll.next();
    timer(0, 60 * 1000)
      .pipe(
        switchMap(() => this.readUserMe()),
        shareReplay({ bufferSize: 1, refCount: true }),
        takeUntil(this.resetPreviousPoll),
      )
      .subscribe();
  }

  readUserMe(): Observable<UserResponse | null> {
    this.isLoading.set(true);
    this.error.set(null);
    return this.api
      .read<UserResponse>('user/me')
      .pipe(
        take(1),
        tap(user => this.loggedInUser.set(user)),
        catchError((httpError: HttpErrorResponse) => {
          this.resetPreviousPoll.next();
          this.error.set(httpError);
          return of(null);
        }),
        finalize(() => this.isLoading.set(false)),
      );
  }

  updateUserBio(requestBody: UpdateUserBioRequestBody): Observable<void> {
    return this.api.update('user/me/bio', requestBody);
  }

  private observeErrorState(): void {
    effect(() => {
      const httpError = this.error();
      untracked(() => {
        if (httpError === null || !this.doesRouteRequireAuth()) return;
        const { error } = httpError;
        const detail = error?.error || error || 'Unknown error';
        this.messageService.add({ severity: 'error', summary: 'Error', detail });
        void this.router.navigate(['/login']);
      });
    });
  }

  private doesRouteRequireAuth(): boolean {
    let route = this.router.routerState.snapshot.root;
    while (route.firstChild) route = route.firstChild;
    return !!route.data['requiresAuth'];
  }
}
