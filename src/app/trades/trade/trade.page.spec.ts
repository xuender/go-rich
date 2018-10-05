import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TradePage } from './trade.page';

describe('TradePage', () => {
  let component: TradePage;
  let fixture: ComponentFixture<TradePage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TradePage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TradePage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
