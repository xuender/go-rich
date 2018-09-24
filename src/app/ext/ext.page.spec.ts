import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ExtPage } from './ext.page';

describe('ExtPage', () => {
  let component: ExtPage;
  let fixture: ComponentFixture<ExtPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ExtPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ExtPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
