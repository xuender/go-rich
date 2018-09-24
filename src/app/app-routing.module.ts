import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', loadChildren: './tabs/tabs.module#TabsPageModule' },
  { path: 'profile', loadChildren: './setting/profile/profile.module#ProfilePageModule' },
  { path: 'items', loadChildren: './items/items.module#ItemsPageModule' },
  { path: 'tags', loadChildren: './tags/tags.module#TagsPageModule' },
  { path: 'tag', loadChildren: './tags/tag/tag.module#TagPageModule' },
  { path: 'select', loadChildren: './tags/select/select.module#SelectPageModule' },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
