import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './auth.guard';
import { LoginComponent } from './login/login.component';
import { EmailComponent } from './email/email.component';
import { HomeComponent } from './home/home.component';
import { RegisterComponent } from './register/register.component';
import { ProjectCreateComponent } from './project-create/project-create.component';
import { ItemCreateComponent } from './item-create/item-create.component';
import { ItemListComponent } from './item-list/item-list.component';
import { ItemShowComponent } from './item-show/item-show.component';

const routes: Routes = [
    { path: '', component: HomeComponent, canActivate: [AuthGuard] },
    { path: 'login', component: LoginComponent },
    { path: 'email', component: EmailComponent },
    { path: 'register', component: RegisterComponent },
    { path: 'projects/new', component: ProjectCreateComponent, canActivate: [AuthGuard] },
    { path: 'items/new', component: ItemCreateComponent, canActivate: [AuthGuard] },
    { path: 'items', component: ItemListComponent, canActivate: [AuthGuard] },
    { path: 'items/:num', component: ItemShowComponent, canActivate: [AuthGuard] },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const routing = RouterModule.forRoot(routes, { useHash: true });
