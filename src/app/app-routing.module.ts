import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', loadChildren: './tabs/tabs.module#TabsPageModule' },
  { path: 'profile', loadChildren: './setting/profile/profile.module#ProfilePageModule' },
  { path: 'items', loadChildren: './items/items.module#ItemsPageModule' },
  { path: 'tags', loadChildren: './tags/tags.module#TagsPageModule' },
  { path: 'tag', loadChildren: './tags/tag/tag.module#TagPageModule' },
  { path: 'select', loadChildren: './tags/select/select.module#SelectPageModule' },
  { path: 'trades', loadChildren: './trades/trades.module#TradesPageModule' },
  { path: 'trade', loadChildren: './trades/trade/trade.module#TradePageModule' },
  { path: 'select', loadChildren: './customers/select/select.module#SelectPageModule' },
  { path: 'select', loadChildren: './items/select/select.module#SelectPageModule' },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
