package com.assingment.android.end_to_end

import androidx.test.rule.ActivityTestRule
import com.agoda.kakao.screen.Screen
import com.assingment.android.screen.HomeScreen
import com.assingment.android.screen.QuizScreen
import com.assingment.android.screen.ResultScreen
import com.assingment.android.view.MainActivity
import org.junit.Rule
import org.junit.Test

class EndToEndHappyPath {
    @get:Rule
    var activityRule: ActivityTestRule<MainActivity> =
        ActivityTestRule(MainActivity::class.java)

    @Test
    fun testE2E() {
        Screen.onScreen<HomeScreen> {
            editText {
                hasText("")
                typeText("Harry Potter")
            }

            startButton.click()
        }

        Screen.onScreen<QuizScreen> {
            submitButton.click()
        }

        Screen.onScreen<ResultScreen> {
            name.hasText("The boy who live")
            score.hasText(0)
            homeButton.click()
        }

        Screen.onScreen<HomeScreen> {
            editText.isDisplayed()
        }

    }

}
