package com.assingment.android.screen

import com.agoda.kakao.screen.Screen
import com.agoda.kakao.text.KButton
import com.assingment.android.R

class QuizScreen : Screen<QuizScreen>() {
    val backButton = KButton { withId(R.id.backButton) }
    val nextButton = KButton { withId(R.id.nextButton) }
    val submitButton = KButton { withId(R.id.nextButton) }
}
