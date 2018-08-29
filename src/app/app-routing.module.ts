import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', loadChildren: './tabs/tabs.module#TabsPageModule' },
  { path: 'customer', loadChildren: './customer/customer.module#CustomerPageModule' },
  { path: 'upload', loadChildren: './contact/upload/upload.module#UploadPageModule' },
  { path: 'group', loadChildren: './contact/group/group.module#GroupPageModule' }
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
