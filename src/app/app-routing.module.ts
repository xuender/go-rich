import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', loadChildren: './tabs/tabs.module#TabsPageModule' },
  { path: 'profile', loadChildren: './setting/profile/profile.module#ProfilePageModule' },
  { path: 'items', loadChildren: './items/items.module#ItemsPageModule' },
  { path: 'tags', loadChildren: './setting/tags/tags.module#TagsPageModule' },
  { path: 'tag', loadChildren: './setting/tags/tag/tag.module#TagPageModule' },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
